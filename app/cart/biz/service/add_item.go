package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hourhl/Qmall/app/cart/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/cart/biz/model"
	"github.com/hourhl/Qmall/app/cart/infra/rpc"
	cart "github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {

	// 1. check user and token
	userVerifyResp, err := rpc.UserClient.VerifyUser(s.ctx, &user.VerifyUserReq{
		UserId: int32(req.UserId),
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}
	if userVerifyResp == nil || userVerifyResp.Res == false {
		return nil, kerrors.NewGRPCBizStatusError(40001, "user verify fail")
	}

	// 2. check product
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product does not exist")
	}

	// 3. add Item
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}

	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{
		Res: true,
	}, nil

}
