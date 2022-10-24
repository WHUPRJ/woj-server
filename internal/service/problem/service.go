package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (*model.Problem, e.Status)
	Update(problem *model.Problem) (*model.Problem, e.Status)
	Query(pid uint, associations bool, shouldEnable bool) (*model.Problem, e.Status)
	QueryFuzz(search string, associations bool, shouldEnable bool) ([]*model.Problem, e.Status)

	CreateVersion(data *CreateVersionData) (*model.ProblemVersion, e.Status)
	UpdateVersion(pvid uint, values interface{}) e.Status
	QueryVersion(pvid uint, shouldEnable bool) (*model.ProblemVersion, e.Status)
	QueryLatestVersion(pid uint) (*model.ProblemVersion, e.Status)
}

type service struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewService(g *global.Global) Service {
	return &service{
		log: g.Log,
		db:  g.Db.Get().(*gorm.DB),
	}
}
