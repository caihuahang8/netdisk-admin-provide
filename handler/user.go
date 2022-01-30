package handler

import (
	"context"
	"github.com/caihuahang8/common"
	"github.com/caihuahang8/netdisk-admin-provide/domain/model"
	"github.com/caihuahang8/netdisk-admin-provide/domain/service"
	user "github.com/caihuahang8/netdisk-admin-provide/proto/admin"
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
