package repository

import (
	"ReserveMate/backend/pkg/domain/model"
)

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	LogIn(email string) (string, error)
}
