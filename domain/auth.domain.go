package domain

import "context"

type AuthService interface {
	Login(c context.Context, username string, password string) (string, error)
	Register(c context.Context, username string, password string) (string, error)
	IsUsernameAlreadyExists(c context.Context, username string) error
}
