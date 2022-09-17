package e

const (
	Success Err = 0
	Unknown Err = 1

	InternalError    Err = 100
	InvalidParameter Err = 101
	NotFound         Err = 102
	DatabaseError    Err = 103

	TokenUnknown   Err = 200
	TokenEmpty     Err = 201
	TokenMalformed Err = 202
	TokenTimeError Err = 203
	TokenInvalid   Err = 204
	TokenSignError Err = 205
	TokenRevoked   Err = 206

	UserNotFound        Err = 300
	UserWrongPassword   Err = 301
	UserDuplicated      Err = 302
	UserUnauthenticated Err = 303

	RedisError Err = 400
)

var msgText = map[Err]string{
	Success: "Success",
	Unknown: "Unknown error",

	InternalError:    "Internal Error",
	InvalidParameter: "Invalid Parameter",
	NotFound:         "Not Found",
	DatabaseError:    "Database Error",

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

	RedisError: "Redis Error",
}
