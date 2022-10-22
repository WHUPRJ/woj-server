package e

const (
	Success Status = iota
	Unknown
)

const (
	InternalError Status = 100 + iota
	InvalidParameter
	NotFound
	DatabaseError
	RedisError
)

const (
	TokenUnknown Status = 200 + iota
	TokenEmpty
	TokenMalformed
	TokenTimeError
	TokenInvalid
	TokenSignError
	TokenRevoked
)

const (
	UserNotFound Status = 300 + iota
	UserWrongPassword
	UserDuplicated
	UserUnauthenticated
	UserUnauthorized
	UserDisabled
)

const (
	ProblemNotFound Status = 500 + iota
	ProblemNotAvailable
	ProblemVersionNotFound
	ProblemVersionNotAvailable
	StatusNotFound
)

const (
	TaskEnqueueFailed Status = 600 + iota
	TaskGetInfoFailed
)

const (
	RunnerDepsBuildFailed Status = 700 + iota
	RunnerDownloadFailed
	RunnerUnzipFailed
	RunnerProblemNotExist
	RunnerProblemPrebuildFailed
	RunnerProblemParseFailed
	RunnerUserNotExist
	RunnerUserCompileFailed
	RunnerRunFailed
	RunnerJudgeFailed
)

var msgText = map[Status]string{
	Success: "Success",
	Unknown: "Unknown error",

	InternalError:    "Internal Error",
	InvalidParameter: "Invalid Parameter",
	NotFound:         "Not Found",
	DatabaseError:    "Database Error",
	RedisError:       "Redis Error",

	TokenUnknown:   "Unknown Error (Token)",
	TokenEmpty:     "Token Empty",
	TokenMalformed: "Token Malformed",
	TokenTimeError: "Token Time Error",
	TokenInvalid:   "Token Invalid",
	TokenSignError: "Token Sign Error",
	TokenRevoked:   "Token Revoked",

	UserNotFound:        "User Not Found",
	UserWrongPassword:   "User Wrong Password",
	UserDuplicated:      "User Duplicated",
	UserUnauthenticated: "User Unauthenticated",
	UserUnauthorized:    "User Unauthorized",
	UserDisabled:        "User Disabled",

	ProblemNotFound:            "Problem Not Found",
	ProblemNotAvailable:        "Problem Not Available",
	ProblemVersionNotFound:     "Problem Version Not Found",
	ProblemVersionNotAvailable: "Problem Version Not Available",

	StatusNotFound: "Status Not Found",

	TaskEnqueueFailed: "Task Enqueue Failed",
	TaskGetInfoFailed: "Task Get Info Failed",

	RunnerDepsBuildFailed:       "Runner Deps Build Failed",
	RunnerDownloadFailed:        "Runner Download Failed",
	RunnerUnzipFailed:           "Runner Unzip Failed",
	RunnerProblemNotExist:       "Runner Problem Not Exist",
	RunnerProblemPrebuildFailed: "Runner Problem Prebuild Failed",
	RunnerProblemParseFailed:    "Runner Problem Parse Failed",
	RunnerUserNotExist:          "Runner User Not Exist",
	RunnerUserCompileFailed:     "Runner User Compile Failed",
	RunnerRunFailed:             "Runner Run Failed",
	RunnerJudgeFailed:           "Runner Judge Failed",
}
