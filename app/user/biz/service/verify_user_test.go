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
		UserId: int32(3),
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzYzOTM5NSwiaWF0IjoxNzMzNjM5Mzk1fQ.359m3KgTdxjHPcCZxfOi51E0waVmjFmaq_JyeRlYcYw",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASSS
}
