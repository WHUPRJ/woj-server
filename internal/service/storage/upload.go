package storage

import (
	"context"
	"github.com/WHUPRJ/woj-server/internal/e"
	"go.uber.org/zap"
	"time"
)

func (s *service) Upload(objectName string, expiry time.Duration) (string, e.Status) {
	preSignedURL, err := s.client.PresignedPutObject(
		context.Background(),
		s.bucket,
		objectName,
		expiry,
	)

	if err != nil {
		s.log.Warn("failed to generate pre-signed upload url",
			zap.Error(err),
			zap.String("objectName", objectName),
			zap.Duration("expiry", expiry),
		)
		return "", e.StorageUploadFailed
	}

	return preSignedURL.String(), e.Success
}
