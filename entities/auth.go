package entity

type Credentials struct {
	Username string
	Password string
}

type CredentialsError struct {
	UsernameError string
	PasswordError string
}

type RegisterCred struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisterCredError struct {
	EmailError           string
	UsernameError        string
	PasswordError        string
	ConfirmPasswordError string
}
