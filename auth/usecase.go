package auth

import (
	"context"
	"github.com/go-park-mail-ru/2021_1_kekEnd/models"
)

type UseCase interface {
	SignUp(ctx context.Context, username, email, firstName, lastName, password string) error

	Login(ctx context.Context, login, password string) bool

	GetUser(ctx context.Context, id int) (*models.User, error)

	UpdateUser(ctx context.Context, id int, newUser *models.User) error
}
