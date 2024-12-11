package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/cart/biz/dal"
	"github.com/hourhl/Qmall/app/cart/infra/rpc"
	cart "github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("load env error : %v\n", err.Error())
	}
	dal.Init()
	rpc.Init()

	ctx := context.Background()
	s := NewAddItemService(ctx)

	// init req and assert value
	reqItem := &cart.CartItem{
		ProductId: uint32(3),
		Quantity:  int32(2),
	}

	req := &cart.AddItemReq{
		UserId: 3,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzY0Mzc3MCwiaWF0IjoxNzMzNjQzNzcwfQ.JleprAcF3_eakTCQEP7o5bceTu5gz40l5BZmIzbmBOM",
		Item:   reqItem,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS

}
