package repository

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user = model.User{Email: email}
	if err := ur.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
