package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/payment/biz/dal"
	payment "github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("loading .env file error : %v\n", err.Error())
	}
	dal.Init()
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		UserId:  uint32(1),
		OrderId: "46f2420c-b545-11ef-b539-28c5c873c72b",
		Amount:  float32(200),
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "4111111111111111",
			CreditCardCvv:             123,
			CreditCardExpirationMonth: 12,
			CreditCardExpirationYear:  2100,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS

}
