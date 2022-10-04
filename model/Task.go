package model

const (
	TypeProblemResolve = "problem:resolve"
	TypeProblemPush    = "problem:push"
	TypeSubmitJudge    = "submit:judge"
)

type ProblemResolvePayload struct {
	ProblemID   uint
	ProblemFile string
}

type ProblemPushPayload struct {
	ProblemFile string
}

type SubmitJudge struct {
	Submission Submission
}
