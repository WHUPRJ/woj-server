package runner

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/down"
	"github.com/WHUPRJ/woj-server/pkg/unzip"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func (s *service) download(version uint, url string) e.Status {
	zipPath := filepath.Join(TmpDir, fmt.Sprintf("%d.zip", version))
	problemPath := filepath.Join(ProblemDir, fmt.Sprintf("%d", version))

	err := down.Down(zipPath, url)
	if err != nil {
		s.log.Error("download problem failed", zap.Error(err))
		return e.RunnerDownloadFailed
	}

	err = unzip.Unzip(zipPath, problemPath)
	if err != nil {
		s.log.Warn("unzip problem failed", zap.Error(err))
		return e.RunnerUnzipFailed
	}

	return e.Success
}

func (s *service) prebuild(version uint, force bool) e.Status {
	if !s.ProblemExists(version) {
		return e.RunnerProblemNotExist
	}

	mark := filepath.Join(ProblemDir, fmt.Sprintf("%d", version), ".mark.prebuild")
	if force {
		_ = os.Remove(mark)
	} else if utils.FileExist(mark) {
		return e.Success
	}

	err := s.execute("problem_prebuild.sh", fmt.Sprintf("%d", version))

	if err != nil {
		s.log.Warn("prebuild problem failed", zap.Error(err), zap.Uint("version", version))
		return e.RunnerProblemPrebuildFailed
	}

	return e.Success
}

func (s *service) NewProblem(version uint, url string, force bool) (Config, e.Status) {
	if force {
		problemPath := filepath.Join(ProblemDir, fmt.Sprintf("%d", version))
		_ = os.RemoveAll(problemPath)
	}

	if !s.ProblemExists(version) {
		status := s.download(version, url)
		if status != e.Success {
			return Config{}, status
		}
	}

	cfg, err := s.ParseConfig(version, false)
	if err != nil {
		return Config{}, e.RunnerProblemParseFailed
	}

	status := s.prebuild(version, true)
	if status != e.Success {
		return Config{}, status
	}

	return cfg, e.Success
}
