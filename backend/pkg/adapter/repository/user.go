package repository

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/middlewares"
	"ReserveMate/backend/pkg/usecase/repository"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	fmt.Println("adapter-repo")
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
	var user model.User
	if result := ur.db.Where("email = ?", email).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) LogIn(email string) (string, error) {
	var user model.User
	if result := ur.db.Where("email = ?", email).First(&user); result.Error != nil {
		return "", fmt.Errorf("user not found")
	}

	t, err := middlewares.GenerateJWTToken(user.Email)
	if err != nil {
		return "", err
	}
	return t, nil
}
