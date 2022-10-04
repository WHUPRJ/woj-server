package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model   `json:"-"`
	SubmissionID uint       `json:"submission_id" gorm:"not null;index"`
	Submission   Submission `json:"-"             gorm:"foreignKey:SubmissionID"`
	Verdict      Verdict    `json:"verdict"       gorm:"not null"`
	Point        int32      `json:"point"         gorm:"not null"`
}
