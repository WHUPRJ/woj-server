package model

type Verdict int

const (
	VerdictJudging Verdict = iota
	VerdictAccepted
	VerdictWrongAnswer
	VerdictTimeLimitExceeded
	VerdictMemoryLimitExceeded
	VerdictRuntimeError
	VerdictCompileError
	VerdictSystemError
	VerdictJuryFailed
	VerdictSkipped
	VerdictPartiallyCorrect
)
