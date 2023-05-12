package repository

import "github.com/benning55/go-clean-arch/pkg/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
