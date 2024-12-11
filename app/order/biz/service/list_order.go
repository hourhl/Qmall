package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hourhl/Qmall/app/order/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/order/biz/model"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	order "github.com/hourhl/Qmall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(7005002, err.Error())
	}
	if list == nil || len(list) == 0 {
		fmt.Printf("list is nil\n")
		return nil, nil
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Street:   v.Consignee.Street,
				City:     v.Consignee.City,
				Province: v.Consignee.Province,
				Country:  v.Consignee.Country,
				ZipCode:  v.Consignee.ZipCode,
			},
			OrderItems: items,
		})
	}

	resp = &order.ListOrderResp{Orders: orders}
	return
}
