package e

// MsgFlags info
var MsgFlags = map[int]string{
	SUCCESS: "ok",
}

// GetMsg INFO
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
