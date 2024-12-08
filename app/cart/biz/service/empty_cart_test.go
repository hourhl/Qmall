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

func TestEmptyCart_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("loading .env files failed: %v\n", err)
	}
	dal.Init()
	rpc.Init()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{
		UserId: 1,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzY0Mzc3MCwiaWF0IjoxNzMzNjQzNzcwfQ.JleprAcF3_eakTCQEP7o5bceTu5gz40l5BZmIzbmBOM",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS
}
