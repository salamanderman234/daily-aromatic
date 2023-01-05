package constanta

import "errors"

var (
	UserDataNotFoundWithCreds = errors.New("Username atau password salah")
	PasswordNotMatch          = errors.New("Password tidak cocok")
	UsernameAlreadyExists     = errors.New("Username sudah digunakan")
)
