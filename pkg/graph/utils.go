package graph

func resolvePagination(first *int, skip *int) (int, int) {
	limit := -1
	offset := -1
	if first != nil {
		limit = *first
	}
	if skip != nil {
		offset = *skip
	}
	return limit, offset
}
