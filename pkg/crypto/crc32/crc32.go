package crc32

import "hash/crc32"

func Encode(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
