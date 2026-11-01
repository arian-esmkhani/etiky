package helper

func HasNext(page, pageSize, total uint8) bool {
	totalPages := (total + pageSize -1) / pageSize
	return page + 1 < totalPages
}