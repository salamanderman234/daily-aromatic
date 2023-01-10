package variable

import "errors"

var (
	// general
	ErrEmptyField   = errors.New("field ini tidak boleh kosong")
	ErrMustNumber   = errors.New("harus berupa angka")
	ErrStringLength = errors.New("harus berisi setidaknya %s karakter dan maksimal %s karakter")
	// database
	ErrDataNotFound = errors.New("data tidak ditemukan")
	// auth
	ErrTokenNotValid = errors.New("sesi ini tidak valid")
	ErrTokenExpired  = errors.New("sesi ini sudah berakhir")
	// user
	ErrMustUniqueUsername   = errors.New("username ini telah digunakan")
	ErrPasswordNotMatch     = errors.New("password tidak sama")
	ErrUserCredsNotFound    = errors.New("username atau password salah")
	ErrPasswordNotQualified = errors.New("password harus mengandung setidaknya 8 karakter dan harus berisi huruf Kapital (A-a) dan juga angka (1-9)")
	ErrBioNotQualified      = errors.New("bio tidak boleh lebih dari 50 karakter")
	// http
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrNotFound       = errors.New("not found")
)
