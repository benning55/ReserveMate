package usecase

import (
	"github.com/benning55/go-clean-arch/pkg/domain/model"
	"github.com/benning55/go-clean-arch/pkg/usecase/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type User interface {
	List(u []*model.User) ([]*model.User, error)
}

func NewUserUsecase(r repository.UserRepository) User {
	return &userUsecase{
		userRepository: r,
	}
}

func (uu *userUsecase) List(u []*model.User) ([]*model.User, error) {
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
