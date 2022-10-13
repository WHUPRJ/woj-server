package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model  `json:"meta"`
	Title       string        `json:"title"        gorm:"not null"`
	Content     string        `json:"content"      gorm:"not null"`
	TimeLimit   uint          `json:"time_limit"   gorm:"not null"`
	MemoryLimit uint          `json:"memory_limit" gorm:"not null"`
	ProviderID  uint          `json:"provider_id"  gorm:"not null;index"`
	Provider    User          `json:"-"            gorm:"foreignKey:ProviderID"`
	Languages   pq.Int32Array `json:"languages"    gorm:"type:int[]"`
	Points      pq.Int32Array `json:"points"       gorm:"type:int[]"`
	StorageKey  string        `json:"storage_key"  gorm:"not null"`
	IsEnabled   bool          `json:"is_enabled"   gorm:"not null;index"`
}
