package convert

import "strconv"

func StringToInt64(number string) int64 {
	value, err := strconv.ParseInt(number, 10, 64)
	if err == nil {
		return value
	}
	return 0
}
func StringToInt32(number string) int64 {
	value, err := strconv.ParseInt(number, 10, 32)
	if err == nil {
		return value
	}
	return 0
}
