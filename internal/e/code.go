package e

const (
	Fallback         Err = 0
	Success          Err = 1
	InternalError    Err = 100
	InvalidParameter Err = 101
	NotFound         Err = 102
)

var msgText = map[Err]string{
	Success:          "Success",
	InternalError:    "Internal Error",
	InvalidParameter: "Invalid Parameter",
	NotFound:         "Not Found",
}
