package constanta

import "errors"

var (
	EmptyCreds   = errors.New("Tidak boleh kosong")
	TokenExpired = errors.New("Sesi ini sudah habis")
	TokenInvalid = errors.New("Sesi ini tidak valid")
	Unauthorized = errors.New("Unauthorized")
)
