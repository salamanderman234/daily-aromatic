package service

const (
	queryLimit = 6
)

func getOffset(page int) int {
	return (page - 1) * queryLimit
}
