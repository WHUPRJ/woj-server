package e

type Status int

func (code Status) String() string {
	msg, ok := msgText[code]
	if ok {
		return msg
	}
	return msgText[InternalError]
}
