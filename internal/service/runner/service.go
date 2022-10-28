package runner

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	// EnsureDeps build docker images
	EnsureDeps(force bool) e.Status
	// NewProblem = Download + Parse + Prebuild
	NewProblem(version uint, url string, force bool) (Config, e.Status)

	// Compile compile user submission
	Compile(version uint, user string, lang string) (JudgeStatus, e.Status)
	// RunAndJudge execute user program
	RunAndJudge(version uint, user string, lang string, config *Config) (JudgeStatus, int32, e.Status)

	// ParseConfig parse config file
	ParseConfig(version uint, skipCheck bool) (Config, error)
	// ProblemExists check if problem exists
	ProblemExists(version uint) bool
}

type service struct {
	log     *zap.Logger
	verbose bool
}

func NewService(g *global.Global) Service {
	return &service{
		log:     g.Log,
		verbose: g.Conf.Development,
	}
}
