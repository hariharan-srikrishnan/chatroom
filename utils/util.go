package utils

var (
	OK       = "OK"
	OK_BYTES = []byte("OK")
)

func StrToBytes(message *string) []byte {
	return []byte(*message)
}

func BytesToStr(bytes []byte) string {
	return string(bytes)
}
