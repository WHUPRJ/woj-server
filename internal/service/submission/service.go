package submission

import (
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/repo/amqp"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
}

type service struct {
	log *zap.Logger
	mq  *amqp.Repo
}

func NewService(g *global.Global) Service {
	return &service{
		log: g.Log,
		mq:  g.Mq.Get().(*amqp.Repo),
	}
}
