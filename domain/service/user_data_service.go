package service

import (
	"errors"
	"github.com/caihuahang8/netdisk-admin-provide/domain/model"
	"github.com/caihuahang8/netdisk-admin-provide/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDateService interface {
	AddUserData(user *model.User) (int64, error)
	CheckPwd(username string, pwd string) (isOk bool, err error)
}

type UserDateService struct {
	UserRepository repository.UserRepository
}

//创建用户
func (u *UserDateService) AddUserData(user *model.User) (int64, error) {
	return u.UserRepository.AddUser(user)
}

//比对账号密码是否正确
func (u *UserDateService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.Password)
}
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}
