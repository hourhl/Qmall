package service

import (
	"context"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {

	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value
	req := &auth.DeliverTokenReq{
		UserId: 3,
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp.Token:\n %s", resp.Token)
	// status : PASS
}
