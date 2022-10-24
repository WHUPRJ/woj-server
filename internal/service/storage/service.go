package storage

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"time"
)

var _ Service = (*service)(nil)

type Service interface {
	Upload(objectName string, expiry time.Duration) (string, e.Status)
	Get(objectName string, expiry time.Duration) (string, e.Status)
}

type service struct {
	log    *zap.Logger
	client *minio.Client
	bucket string
}

func NewService(g *global.Global) Service {
	minioClient, err := minio.New(g.Conf.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(g.Conf.Storage.AccessKey, g.Conf.Storage.SecretKey, ""),
		Secure: g.Conf.Storage.UseSSL,
	})

	if err != nil {
		g.Log.Fatal("failed to create minio client", zap.Error(err))
		return nil
	}

	return &service{
		log:    g.Log,
		client: minioClient,
		bucket: g.Conf.Storage.Bucket,
	}
}
