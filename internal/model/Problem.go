package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model `json:"meta"`
	Title      string `json:"title"      gorm:"not null"`
	Statement  string `json:"statement"  gorm:"not null"`
	ProviderID uint   `json:"-"          gorm:"not null;index"`
	Provider   User   `json:"provider"   gorm:"foreignKey:ProviderID"`
	IsEnabled  bool   `json:"is_enabled" gorm:"not null;index"`
}

type ProblemVersion struct {
	gorm.Model `json:"meta"`
	ProblemID  uint        `json:"-"          gorm:"not null;index"`
	Context    pgtype.JSON `json:"context"    gorm:"type:json"`
	StorageKey string      `json:"-"          gorm:"not null"`
	IsEnabled  bool        `json:"is_enabled" gorm:"not null;index"`
}
