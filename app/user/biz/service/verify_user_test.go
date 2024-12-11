package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/infra/rpc"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestVerifyUser_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	mysql.Init()
	rpc.Init()
	ctx := context.Background()
	s := NewVerifyUserService(ctx)
	if err != nil {
		fmt.Printf("loading .env files failed: %s\n", err.Error())
		return
	}
	// init req and assert value

	req := &user.VerifyUserReq{
		UserId: 3,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzkwMzUzMiwiaWF0IjoxNzMzOTAzNTMyfQ.87f9phamxfBfhI0xnTZSdFKqHCugTl0V3BTmIRX25mI",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASSS
}
