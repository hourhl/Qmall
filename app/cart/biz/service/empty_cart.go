package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hourhl/Qmall/app/cart/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/cart/biz/model"
	"github.com/hourhl/Qmall/app/cart/infra/rpc"
	cart "github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {

	// verify user
	if req == nil {
		return nil, kerrors.NewGRPCBizStatusError(50001, err.Error())
	}
	verifyResp, err := rpc.UserClient.VerifyUser(s.ctx, &user.VerifyUserReq{
		UserId: int32(req.UserId),
		Token:  req.Token,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	if verifyResp == nil || verifyResp.Res == false {
		return nil, kerrors.NewBizStatusError(50003, "user verify fail")
	}

	// Empty cart
	err = model.EmptyCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50003, err.Error())
	}

	return &cart.EmptyCartResp{}, nil
}
