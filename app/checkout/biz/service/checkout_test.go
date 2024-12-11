package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/checkout/biz/dal"
	"github.com/hourhl/Qmall/app/checkout/infra/rpc"
	checkout "github.com/hourhl/Qmall/rpc_gen/kitex_gen/checkout"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("loading .env file error : %v\n", err.Error())
	}
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	dal.Init()
	rpc.Init()
	// init req and assert value
	address := &checkout.Address{
		StreetAddress: "zhongguancunstreet",
		City:          "Beijing",
		Province:      "BJ",
		Country:       "China",
		ZipCode:       "10001",
	}

	creditcard := &payment.CreditCardInfo{
		CreditCardNumber:          "4111111111111111",
		CreditCardCvv:             123,
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  2100,
	}

	req := &checkout.CheckoutReq{
		UserId:     uint32(3),
		Email:      "hourhl@hourl.com",
		Address:    address,
		CreditCard: creditcard,
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzkwMzUzMiwiaWF0IjoxNzMzOTAzNTMyfQ.87f9phamxfBfhI0xnTZSdFKqHCugTl0V3BTmIRX25mI",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS without verify user in cart module

}
