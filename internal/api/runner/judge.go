package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/runner"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

func (h *handler) Judge(_ context.Context, t *asynq.Task) error {
	var p model.SubmitJudgePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	user := utils.RandomString(16)
	h.log.Info("judge", zap.Any("payload", p), zap.String("user", user))

	status, point, ctx := func() (e.Status, int32, runner.JudgeStatus) {
		systemError := runner.JudgeStatus{Message: "System Error"}

		// 1. write user code
		userCode := filepath.Join(runner.UserDir, user, fmt.Sprintf("%s.%s", user, p.Submission.Language))
		if !utils.FileTouch(userCode) {
			return e.InternalError, 0, systemError
		}
		err := utils.FileWrite(userCode, []byte(p.Submission.Code))
		if err != nil {
			return e.InternalError, 0, systemError
		}

		// 2. check problem
		if !h.runnerService.ProblemExists(p.ProblemVersionID) {
			url, status := h.storageService.Get(p.StorageKey, time.Second*60*5)
			if status != e.Success {
				return e.InternalError, 0, systemError
			}

			_, status = h.runnerService.NewProblem(p.ProblemVersionID, url, false)
			if status != e.Success {
				return e.InternalError, 0, systemError
			}
		}

		// 3. compile
		compileResult, status := h.runnerService.Compile(p.ProblemVersionID, user, p.Submission.Language)
		if status != e.Success {
			return e.Success, 0, compileResult
		}

		// 4. config
		config, err := h.runnerService.ParseConfig(p.ProblemVersionID, true)
		if err != nil {
			return e.InternalError, 0, systemError
		}

		// 5. run and judge
		result, point, status := h.runnerService.RunAndJudge(p.ProblemVersionID, user, p.Submission.Language, &config)
		return utils.If(status != e.Success, e.InternalError, e.Success), point, result
	}()

	h.taskService.SubmitUpdate(&model.SubmitUpdatePayload{
		Status:           status,
		SubmissionID:     p.Submission.ID,
		ProblemVersionID: p.ProblemVersionID,
		Point:            point,
	}, ctx)

	return nil
}
