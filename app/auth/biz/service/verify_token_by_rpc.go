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
	fmt.Printf("tokenString : %s\n", tokenString)
	fmt.Printf("Verify by cache\n")
	cachedKey := fmt.Sprintf("%s_%d", "token", req.UserId)
	cachedResult, err := redis.RedisClient.Get(s.ctx, cachedKey).Result()
	if err != nil {
		fmt.Sprintf("Fail to get token from cache\n")
		return &auth.VerifyResp{Res: false}, err
	}
	if cachedResult != tokenString {
		fmt.Sprintf("token not match\n")
		return &auth.VerifyResp{Res: false}, nil
	}

	fmt.Printf("Congratulation!!!\n")
	resp = &auth.VerifyResp{Res: true}
	return resp, nil
}
