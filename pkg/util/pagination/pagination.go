package pagination

type Query interface {
	Offset(offset int32)
	Limit(limit int32)
}

func GetPageOffset(pageNum, pageSize int32) int {
	return int((pageNum - 1) * pageSize)
}

func GetTotalPages(total, pageSize int32) int32 {
	return total / pageSize
}
