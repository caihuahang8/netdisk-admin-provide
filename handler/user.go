package handler

import (
	"context"
	"github.com/caihuahang8/admin/domain/model"
	"github.com/caihuahang8/admin/domain/service"
	"github.com/caihuahang8/admin/proto/user"
	"github.com/caihuahang8/common"
)

type Service struct {
	UserDataService service.IUserDateService
}

func (u *Service) AddUser(ctx context.Context, request *user.UserAdd, response *user.Response) (err error) {
	user := &model.User{}
	err = common.SwapTo(user, request)
	_, err = u.UserDataService.AddUserData(user)
	return err
}
