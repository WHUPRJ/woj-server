package model

import (
	"github.com/WHUPRJ/woj-server/internal/e"
)

const (
	TypeProblemBuild  = "problem:build"
	TypeProblemUpdate = "problem:update"
	TypeSubmitJudge   = "submit:judge"
	TypeSubmitUpdate  = "submit:update"
)

const (
	QueueServer = "server"
	QueueRunner = "runner"
)

type ProblemBuildPayload struct {
	ProblemVersionID uint
	StorageKey       string
}

type ProblemUpdatePayload struct {
	Status           e.Status
	ProblemVersionID uint
	Context          string
}

type SubmitJudgePayload struct {
	ProblemVersionID uint
	StorageKey       string
	Submission       Submission
}

type SubmitUpdatePayload struct {
	Status           e.Status
	SubmissionID     uint
	ProblemVersionID uint
	Point            int32
	Context          string
}
