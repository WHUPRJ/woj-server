package model

type Verdict int

const (
	VerdictJudging             Verdict = 1
	VerdictAccepted            Verdict = 2
	VerdictWrongAnswer         Verdict = 3
	VerdictTimeLimitExceeded   Verdict = 4
	VerdictMemoryLimitExceeded Verdict = 5
	VerdictRuntimeError        Verdict = 6
	VerdictCompileError        Verdict = 7
	VerdictSystemError         Verdict = 8
	VerdictJuryFailed          Verdict = 9
	VerdictSkipped             Verdict = 10
	VerdictPartiallyCorrect    Verdict = 11
)
