package pagination

type Query interface {
	Offset(offset int)
	Limit(limit int)
}

func GetPageOffset(pageNum, pageSize int64) int64 {
	return (pageNum - 1) * pageSize
}
