package e

const (
	Success Status = 0
	Unknown Status = 1

	InternalError    Status = 100
	InvalidParameter Status = 101
	NotFound         Status = 102
	DatabaseError    Status = 103
	RedisError       Status = 104

	TokenUnknown   Status = 200
	TokenEmpty     Status = 201
	TokenMalformed Status = 202
	TokenTimeError Status = 203
	TokenInvalid   Status = 204
	TokenSignError Status = 205
	TokenRevoked   Status = 206

	UserNotFound        Status = 300
	UserWrongPassword   Status = 301
	UserDuplicated      Status = 302
	UserUnauthenticated Status = 303
	UserUnauthorized    Status = 304
	UserDisabled        Status = 305

	ProblemNotFound Status = 500
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

	ProblemNotFound: "Problem Not Found",
}
