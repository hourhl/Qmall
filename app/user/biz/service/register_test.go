package service

import (
	"context"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "hourhl1@hourhl.com",
		Password:        "123hwih",
		ConfirmPassword: "123hwih",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
