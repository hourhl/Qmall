package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/order/biz/dal"
	order "github.com/hourhl/Qmall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestListOrder_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("loading .env file failed: %v\n", err)
	}
	dal.Init()
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{
		UserId: uint32(3),
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	// status : PASS

}
