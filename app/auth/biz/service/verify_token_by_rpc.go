package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/auth/biz/dal/redis"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	tokenString := req.Token
	cachedKey := fmt.Sprintf("%s_%d", "token", req.UserId)
	cachedResult, err := redis.RedisClient.Get(s.ctx, cachedKey).Result()
	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}
	if cachedResult != tokenString {
		return &auth.VerifyResp{Res: false}, nil
	}
	resp = &auth.VerifyResp{Res: true}
	return resp, nil
}
