package model

const (
	TypeProblemPush = "problem:push"
	TypeSubmitJudge = "submit:judge"
)

type ProblemPushPayload struct {
	ProblemID   uint
	ProblemFile string
}

type SubmitJudge struct {
	Submission Submission
}
