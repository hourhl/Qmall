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
	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhvdXJobCIsIm5iZiI6MTczMzM2ODkwMywiaWF0IjoxNzMzMzY4OTAzfQ.oYM-5_RONxu4BYEGcngEUDIEw-SjbTEnaz_4i2kIhFs",
	}
	resp, err := s.Run(req)
	t.Logf("err:%v", err)
	t.Logf("resp:%v", resp)

	// wrong token
	req = &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImlzcyI6ImhobCIsIm5iZiI6MTczMzM2ODkwMywiaWF0IjoxNzMzMzY4OTAzfQ.n6MCk1Znjfio8AhZu7e2vZxqmmci9Ey8rXTzC-C3Hxw",
	}
	resp, err = s.Run(req)
	t.Logf("err:%v", err)
	t.Logf("resp:%v", resp)

	// status : PASS
}
