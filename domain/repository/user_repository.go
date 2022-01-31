package repository

import (
	"errors"
	"github.com/caihuahang8/netdisk-admin-provide/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	AddUser(user *model.User) (int64, error)
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
}

//初始化
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func (u *UserRepository) AddUser(user *model.User) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(user, model.User{ID: user.ID, Password: user.Password, Username: user.Username})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("添加用户失败")
	}
	return user.ID, nil
}

//根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("username = ?", name).Find(user).Error
}

//初始化数据库表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}
