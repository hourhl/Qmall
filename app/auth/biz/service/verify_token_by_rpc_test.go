package service

import (
	"context"
	"github.com/hourhl/Qmall/app/auth/biz/dal"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value
	req := &auth.VerifyTokenReq{
		UserId: 3,
		Token:  "JIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczNDQ5MDE2MywiaWF0IjoxNzM0NDkwMTYzfQ.hC_8IT0l5wsMV4cNLmebke6GYnrDFqHKlt69cNYdfsQ",
	}
	resp, err := s.Run(req)
	t.Logf("err:%v", err)
	t.Logf("resp:%v", resp)

	// status : PASS
}
