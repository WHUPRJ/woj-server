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
)

func (h *handler) Judge(_ context.Context, t *asynq.Task) error {
	var p model.SubmitJudgePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	user := utils.RandomString(16)
	h.log.Info("judge", zap.Any("payload", p), zap.String("user", user))

	// common
	systemError := runner.JudgeStatus{Message: "System Error"}

	// write code
	userCode := filepath.Join(runner.UserDir, user, fmt.Sprintf("%s.%s", user, p.Submission.Language))
	if !utils.FileTouch(userCode) {
		h.log.Info("Touch file failed", zap.String("userCode", userCode))
		h.taskService.SubmitUpdate(e.InternalError, p.ProblemVersionId, 0, systemError)
		return nil
	}
	err := utils.FileWrite(userCode, []byte(p.Submission.Code))
	if err != nil {
		h.log.Info("Write file failed", zap.String("code", p.Submission.Code))
		h.taskService.SubmitUpdate(e.InternalError, p.ProblemVersionId, 0, systemError)
		return nil
	}

	// compile
	result, status := h.runnerService.Compile(p.ProblemVersionId, user, p.Submission.Language)
	if status == e.RunnerProblemNotExist {
		_, status := h.runnerService.NewProblem(p.ProblemVersionId, p.StorageKey)
		if status != e.Success {
			h.log.Warn("download problem failed",
				zap.Any("status", status),
				zap.Uint("pvid", p.ProblemVersionId),
				zap.String("storageKey", p.StorageKey))
			h.taskService.SubmitUpdate(status, p.ProblemVersionId, 0, systemError)
			return nil
		}
	} else if status != e.Success {
		h.taskService.SubmitUpdate(status, p.Submission.ID, 0, result)
		return nil
	}

	// config
	config, err := h.runnerService.ParseConfig(p.ProblemVersionId, true)
	if err != nil {
		h.log.Info("parse config failed", zap.Error(err), zap.Uint("pvid", p.ProblemVersionId))
		h.taskService.SubmitUpdate(e.InternalError, p.ProblemVersionId, 0, systemError)
		return nil
	}

	// run
	var points int32
	result, points, status = h.runnerService.RunAndJudge(p.ProblemVersionId, user, p.Submission.Language, &config)
	h.taskService.SubmitUpdate(status, p.Submission.ID, points, result)

	return nil
}
