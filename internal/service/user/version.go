package user

import (
	"context"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"go.uber.org/zap"
)

func (s *service) IncrVersion(id uint) (int64, e.Status) {
	version, err := s.redis.Incr(context.Background(), fmt.Sprintf("Version:%d", id)).Result()
	if err != nil {
		s.log.Debug("redis.Incr error", zap.Error(err))
		return -1, e.RedisError
	}
	return version, e.Success
}
