package service

import (
	"context"
	"errors"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/biz/model"
	"github.com/hourhl/Qmall/app/user/infra/rpc"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
)

type VerifyUserService struct {
	ctx context.Context
} // NewVerifyUserService new VerifyUserService
func NewVerifyUserService(ctx context.Context) *VerifyUserService {
	return &VerifyUserService{ctx: ctx}
}

// Run create note info
func (s *VerifyUserService) Run(req *user.VerifyUserReq) (resp *user.VerifyUserResp, err error) {
	// 1. check user
	if req == nil {
		return nil, errors.New("user id is nil")
	}
	row, err := model.GetById(mysql.DB, int(req.UserId))
	if row == nil || err != nil {
		return nil, errors.New("user not exist")
	}

	// 2. check token
	if rpc.AuthClient == nil {
		return nil, errors.New("auth client is nil")
	}
	tokenVerify, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if tokenVerify == nil || err != nil {
		return nil, errors.New("verify token error")
	}

	resp = &user.VerifyUserResp{
		Res: tokenVerify.Res,
	}
	return resp, err
}
