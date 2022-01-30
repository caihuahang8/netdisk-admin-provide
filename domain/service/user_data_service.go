package service

import (
	"github.com/caihuahang8/admin/domain/model"
	"github.com/caihuahang8/admin/domain/repository"
)

type IUserDateService interface {
	AddUserData(user *model.User) (int64, error)
}

type UserDateService struct {
	UserRepository repository.UserRepository
}

//创建用户
func (u *UserDateService) AddUserData(user *model.User) (int64, error) {
	return u.UserRepository.AddUser(user)
}
