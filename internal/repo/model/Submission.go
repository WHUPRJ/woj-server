package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model `json:"-"`
	ProblemID  uint    `json:"problem_id" gorm:"not null;index"`
	Problem    Problem `json:"-"          gorm:"foreignKey:ProblemID"`
	UserID     uint    `json:"user_id"    gorm:"not null;index"`
	User       User    `json:"-"          gorm:"foreignKey:UserID"`
	Language   Lang    `json:"language"   gorm:"not null"`
	Code       string  `json:"code"       gorm:"not null"`
}
