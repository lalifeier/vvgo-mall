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

// int64 to string
func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}
