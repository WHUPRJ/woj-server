package global

import (
	"gorm.io/gorm"
)

type Repo interface {
	Setup(*Global)

	Get() *gorm.DB
	Close() error
}
