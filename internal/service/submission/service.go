package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(data *CreateData) (*model.Submission, e.Status)
	Query(pid uint, uid uint, offset int, limit int) ([]*model.Submission, e.Status)
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
