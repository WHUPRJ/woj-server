package runner

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"os"
	"path/filepath"
)

func (s *service) Compile(version uint, user string, lang string) (JudgeStatus, e.Status) {
	target := filepath.Join(UserDir, user, fmt.Sprintf("%s.out", user))

	_ = os.Remove(target)
	status := s.checkAndExecute(version, user, lang, "problem_compile.sh", e.RunnerUserCompileFailed)

	log := filepath.Join(UserDir, user, fmt.Sprintf("%s.compile.log", user))
	msg, err := utils.FileRead(log)
	msg = utils.If(err == nil, msg, nil)
	msgText := string(msg)

	if !utils.FileExist(target) || utils.FileEmpty(target) {
		return JudgeStatus{
				Message: "compile failed",
				Tasks:   []TaskStatus{{Verdict: VerdictCompileError, Message: msgText}}},
			utils.If(status == e.Success, e.RunnerUserCompileFailed, status)
	}

	return JudgeStatus{}, e.Success
}
