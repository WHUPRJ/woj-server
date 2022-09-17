package e

const (
	Success          Err = 0
	Unknown          Err = 1
	InternalError    Err = 100
	InvalidParameter Err = 101
	NotFound         Err = 102
	DatabaseError    Err = 103
)

var msgText = map[Err]string{
	Success:          "Success",
	Unknown:          "Unknown error",
	InternalError:    "Internal Error",
	InvalidParameter: "Invalid Parameter",
	NotFound:         "Not Found",
	DatabaseError:    "Database Error",
}
