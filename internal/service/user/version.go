package user

import (
	"context"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"go.uber.org/zap"
)

func (s *service) IncrVersion(uid uint) (int64, e.Status) {
	version, err := s.redis.Incr(context.Background(), fmt.Sprintf("Version:%d", uid)).Result()
	if err != nil {
		s.log.Warn("RedisError", zap.Error(err), zap.Any("uid", uid))
		return -1, e.RedisError
	}

	return version, e.Success
}
