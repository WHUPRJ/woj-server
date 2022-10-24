package model

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model       `json:"meta"`
	SubmissionID     uint        `json:"-"                  gorm:"not null;index"`
	Submission       Submission  `json:"submission"         gorm:"foreignKey:SubmissionID"`
	ProblemVersionID uint        `json:"problem_version_id" gorm:"not null;index"`
	Context          pgtype.JSON `json:"context"            gorm:"type:json;not null"`
	Point            int32       `json:"point"              gorm:"not null"`
	IsEnabled        bool        `json:"is_enabled"         gorm:"not null;index"`
}
