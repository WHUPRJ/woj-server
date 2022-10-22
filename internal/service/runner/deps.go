package runner

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *service) EnsureDeps(force bool) e.Status {
	mark := filepath.Join(Prefix, ".mark.docker")

	if force {
		_ = os.Remove(mark)
	} else if utils.FileExist(mark) {
		return e.Success
	}

	script := filepath.Join(ScriptsDir, "prepare_container.sh")
	cmd := exec.Command(script)
	cmd.Dir = ScriptsDir
	err := cmd.Run()
	if err != nil {
		s.log.Warn("prebuild docker images failed", zap.Error(err))
		return e.RunnerDepsBuildFailed
	}

	return e.Success
}
