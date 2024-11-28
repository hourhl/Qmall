package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/hourhl/Qmall/app/order/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/order/biz/model"
	order "github.com/hourhl/Qmall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		err = kerrors.NewGRPCBizStatusError(7005001, "order item is empty")
		return nil, err
	}
	// 因为涉及到两个表的操作，因此需要用事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			o.Consignee.Street = req.Address.Street
			o.Consignee.City = req.Address.City
			o.Consignee.Province = req.Address.Province
			o.Consignee.Country = req.Address.Country
			o.Consignee.ZipCode = req.Address.ZipCode
		}
		if err := tx.Model(&model.Order{}).Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				ProductId:    v.Item.ProductId,
				OrderIdRefer: orderId.String(),
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}
		if err := tx.Model(&model.OrderItem{}).Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{OrderId: orderId.String()},
		}

		return nil
	})
	return
}
