package runner

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	Prefix     = "./resource/runner"
	ProblemDir = "./problem/"
	ScriptsDir = "./scripts/"
	UserDir    = "./user/"
	TmpDir     = "./tmp/"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	Prefix = path.Join(wd, Prefix)
	ProblemDir = path.Join(Prefix, ProblemDir)
	ScriptsDir = path.Join(Prefix, ScriptsDir)
	UserDir = path.Join(Prefix, UserDir)
	TmpDir = path.Join(Prefix, TmpDir)
}

func (s *service) execute(script string, args ...string) error {
	p := filepath.Join(ScriptsDir, script)
	cmd := exec.Command(p, args...)
	cmd.Dir = ScriptsDir
	if s.verbose {
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

func (s *service) checkAndExecute(version uint, user string, lang string, script string, fail e.Status) e.Status {
	if !s.ProblemExists(version) {
		s.log.Info("problem not exists", zap.Uint("version", version))
		return e.RunnerProblemNotExist
	}

	if !s.userExists(user, fmt.Sprintf("%s.%s", user, lang)) {
		s.log.Info("user program not exists", zap.String("user", user), zap.String("lang", lang))
		return e.RunnerUserNotExist
	}

	err := s.execute(script, fmt.Sprintf("%d", version), user, lang)

	if err != nil {
		s.log.Info("execute failed",
			zap.Error(err),
			zap.Uint("version", version),
			zap.String("user", user),
			zap.String("lang", lang))
		return fail
	}

	return e.Success
}

func (s *service) ProblemExists(version uint) bool {
	problemPath := filepath.Join(ProblemDir, fmt.Sprintf("%d", version))
	return utils.FileExist(problemPath)
}

func (s *service) userExists(user string, file string) bool {
	userPath := filepath.Join(UserDir, user, file)
	return utils.FileExist(userPath)
}
