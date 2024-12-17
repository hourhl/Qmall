package service

import (
	"context"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	// right token
	//req := &auth.VerifyTokenReq{
	//	UserId: 1,
	//	Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzYzODA2NywiaWF0IjoxNzMzNjM4MDY3fQ.OCXcnu9Qtim9H19qxVBKUpASgahu6gprXFwQRX6seN4",
	//}
	//resp, err := s.Run(req)
	//t.Logf("err:%v", err)
	//t.Logf("resp:%v", resp)

	// wrong token
	req := &auth.VerifyTokenReq{
		UserId: 3,
		Token:  "IUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczNDQwNzIyOCwiaWF0IjoxNzM0NDA3MjI4fQ.kbx8OvLSEV8KiHRJLGR6y84rxbQ-nSqdct-Ei0V5tvU",
	}
	resp, err := s.Run(req)
	t.Logf("err:%v", err)
	t.Logf("resp:%v", resp)

	// status : PASS
}
