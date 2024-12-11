package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/order/biz/dal"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	order "github.com/hourhl/Qmall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("loading .env file failed : %s\n", err.Error())
	}
	dal.Init()
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	item1 := &cart.CartItem{
		ProductId: 1,
		Quantity:  10,
	}
	cost1 := 99
	orderItem1 := &order.OrderItem{
		Item: item1,
		Cost: float32(cost1),
	}

	item2 := &cart.CartItem{
		ProductId: 2,
		Quantity:  2,
	}
	cost2 := 17.6
	orderItem2 := &order.OrderItem{
		Item: item2,
		Cost: float32(cost2),
	}

	orderItems := []*order.OrderItem{orderItem1, orderItem2}

	req := &order.PlaceOrderReq{
		UserId:       uint32(1),
		UserCurrency: "RMB",
		Address: &order.Address{
			Country:  "CN",
			Province: "BJ",
			City:     "BEIJING",
			Street:   "haidian",
			ZipCode:  int32(100000),
		},
		Email:      "hourhl3@hourhl.com",
		OrderItems: orderItems,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS

}
