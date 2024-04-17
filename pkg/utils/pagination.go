package utils

// GetPage get page parameters
func GetPage(page, size int) (skip int, limit int) {
	limit = size
	if limit <= 0 {
		limit = 0
	}
	if page > 0 {
		skip = (page - 1) * limit
	}

	return skip, limit
}
