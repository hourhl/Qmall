package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hourhl/Qmall/app/cart/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/cart/biz/model"
	"github.com/hourhl/Qmall/app/cart/infra/rpc"
	cart "github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// verify user
	if req == nil {
		fmt.Printf("user do not exist")
		return nil, nil
	}
	verifyResp, err := rpc.UserClient.VerifyUser(s.ctx, &user.VerifyUserReq{
		UserId: int32(req.UserId),
		Token:  req.Token,
	})
	if err != nil {
		fmt.Printf("verify user err: %v\n", err)
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	if verifyResp == nil || verifyResp.Res == false {
		fmt.Printf("user verify fail\n")
		return nil, kerrors.NewBizStatusError(50002, "user do not exist")
	}

	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}

	var items []*cart.CartItem
	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Qty),
		})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.UserId, Items: items}}, nil
}
