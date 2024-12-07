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

func TestLogin_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("godotenv.Load %v", err)
	}
	mysql.Init()
	rpc.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value
	req := &user.LoginReq{
		Email:    "hourhl1@hourhl.com",
		Password: "123hwih",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
