package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(uint, *model.Problem) (*model.Problem, e.Status)
	Update(uint, *model.Problem) (*model.Problem, e.Status)
	Query(uint) (*model.Problem, e.Status)
	QueryFuzz(string) ([]*model.Problem, e.Status)
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
