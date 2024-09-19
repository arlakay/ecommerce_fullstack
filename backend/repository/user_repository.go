package repository

import (
	"backend_ersa/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(user model.User) (uint, error)
	Register(user *model.User) error
	UsersAll() ([]model.User, error)
	UserByID(id uint) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Register(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) Login(user model.User) (uint, error) {
	var existingUser model.User

	if err := u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, err
		}
		return 0, err
	}

	return existingUser.ID, nil
}

func (u *userRepository) UsersAll() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) UserByID(id uint) (model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil
		} else {
			return model.User{}, err
		}
	}
	return user, nil

}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil
		} else {
			return model.User{}, err
		}
	}
	return user, nil
}
