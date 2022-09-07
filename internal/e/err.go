package e

type Err int

func (code Err) String() string {
	msg, ok := msgText[code]
	if ok {
		return msg
	}
	return msgText[InternalError]
}
