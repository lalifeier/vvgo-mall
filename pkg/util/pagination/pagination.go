package pagination

type Query interface {
	Offset(offset int64)
	Limit(limit int64)
}

func GetPageOffset(pageNum, pageSize int64) int64 {
	return (pageNum - 1) * pageSize
}
