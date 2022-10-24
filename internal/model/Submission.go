package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model `json:"meta"`
	ProblemID  uint   `json:"problem_id" gorm:"not null;index"`
	UserID     uint   `json:"-"          gorm:"not null;index"`
	User       User   `json:"user"       gorm:"foreignKey:UserID"`
	Language   string `json:"language"   gorm:"not null"`
	Code       string `json:"code"       gorm:"not null"`
}
