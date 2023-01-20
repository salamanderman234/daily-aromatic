package variable

const (
	TokenAndSessionValidityRange = 24
	QueryLimit                   = 6
	SearchLimit                  = 8
)

func DefaultOffset(page int) int {
	return (page - 1) * QueryLimit
}
func SearchOffset(page int) int {
	return (page - 1) * SearchLimit
}
