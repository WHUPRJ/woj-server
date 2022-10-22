package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(*model.Status) (*model.Status, e.Status)
	Query(sid uint, associations bool) (*model.Status, e.Status)
	QueryByVersion(pvid uint, offset int, limit int) ([]*model.Status, e.Status)
	Rejudge(statusID uint) ([]*model.Status, e.Status)
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
