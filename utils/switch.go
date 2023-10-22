package utils

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func ByteToBool(b byte) bool {
	return b != 0
}
func IntToBool(i int) bool {
	return i != 0
}
