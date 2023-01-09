package constanta

import "errors"

var (
	BadRequestError     = errors.New("Bad Request")
	InternalServerError = errors.New("Internal Server Error")
	NotFoundError       = errors.New("Not Found")
)
