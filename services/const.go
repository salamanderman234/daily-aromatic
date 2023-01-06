package service

const (
	queryLimit  = 6
	searchLimit = 8
)

func getOffset(page int) int {
	return (page - 1) * queryLimit
}

func getSarchOffset(page int) int {
	return (page - 1) * searchLimit
}
