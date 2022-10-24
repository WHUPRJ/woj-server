package storage

import (
	"context"
	"github.com/WHUPRJ/woj-server/internal/e"
	"go.uber.org/zap"
	"net/url"
	"time"
)

func (s *service) Get(objectName string, expiry time.Duration) (string, e.Status) {
	preSignedURL, err := s.client.PresignedGetObject(
		context.Background(),
		s.bucket,
		objectName,
		expiry,
		url.Values{},
	)

	if err != nil {
		s.log.Warn("failed to generate pre-signed get url",
			zap.Error(err),
			zap.String("objectName", objectName),
			zap.Duration("expiry", expiry),
		)
		return "", e.StorageGetFailed
	}

	return preSignedURL.String(), e.Success
}
