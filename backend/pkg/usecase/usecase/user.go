package usecase

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/repository"
	"errors"
	"fmt"
)

type userUsecase struct {
	userRepository repository.UserRepository
	dBRepository   repository.DBRepository
}

type User interface {
	List(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	LogIn(email string) (string, error)
}

func NewUserUsecase(r repository.UserRepository, d repository.DBRepository) User {
	return &userUsecase{
		userRepository: r,
		dBRepository:   d,
	}
}

func (uu *userUsecase) List(u []*model.User) ([]*model.User, error) {
	fmt.Println("use-case")
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uu *userUsecase) Create(u *model.User) (*model.User, error) {
	_, err := uu.GetUserByEmail(u.Email)
	if err == nil {
		return nil, fmt.Errorf("user with this email already exists")
	}
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.userRepository.Create(u)

		return u, err
	})
	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) GetUserByEmail(email string) (*model.User, error) {
	u, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uu *userUsecase) LogIn(email string) (string, error) {
	token, err := uu.userRepository.LogIn(email)
	if err != nil {
		return "", err
	}
	return token, nil
}
