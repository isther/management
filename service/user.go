package service

import (
	"errors"

	"github.com/isther/management/dao"
	"github.com/isther/management/model"
)

type UserService struct{}

func NewUserService() *UserService { return &UserService{} }

func (service *UserService) Create(user model.User) error {
	if user.Username == "" {
		return errors.New("username is empty")
	}

	tx := dao.DB.Create(&model.UserSql{User: user})
	return tx.Error
}

func (service *UserService) FindUserByUsername(username string) (model.UserSql, error) {
	var (
		userSql model.UserSql
	)
	tx := dao.DB.Where("username = ?", username).First(&userSql)
	return userSql, tx.Error
}
