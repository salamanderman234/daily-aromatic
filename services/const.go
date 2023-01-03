package service

const (
	productPerPage = 6
)

func getOffset(page int) int {
	return (page - 1) * productPerPage
}
