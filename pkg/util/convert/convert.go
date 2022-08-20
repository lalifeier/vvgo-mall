package convert

import (
	"bytes"
	"io/ioutil"
	"strconv"
)

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

func FileToReader(path string) (*bytes.Reader, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	return reader, nil
}
