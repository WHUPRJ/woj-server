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
	ProblemFile      string
}

type ProblemUpdatePayload struct {
	Status           e.Status
	ProblemVersionID uint
	Context          string
}

type SubmitJudgePayload struct {
	ProblemVersionId uint
	StorageKey       string
	Submission       Submission
}

type SubmitUpdatePayload struct {
	Status  e.Status
	Sid     uint
	Point   int32
	Context string
}
