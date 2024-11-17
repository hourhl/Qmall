package service

import (
	"context"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

// TODO
func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "hourhl@hourhl.com",
		Password:        "123hwih",
		ConfirmPassword: "123hwih",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
