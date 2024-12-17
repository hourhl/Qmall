package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hourhl/Qmall/app/auth/biz/dal/redis"
	"github.com/hourhl/Qmall/app/auth/biz/model"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"os"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.

	fmt.Printf("now is deliver token and the user id is %d\n", req.UserId)
	claim := &model.Claim{
		req.UserId,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			//不设置 ExpiresAt， 默认不过期
			Issuer: "hourhl",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// unit test
	//err = godotenv.Load("../../.env")

	if err != nil {
		kerrors.NewGRPCBizStatusError(1001, "Get env error")
		return nil, err
	}
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	ss, err := token.SignedString(mySigningKey)
	resp = &auth.DeliveryResp{Token: ss}

	// Redis: Cache
	cachedKey := fmt.Sprintf("%s_%d", "token", req.UserId)
	if err != nil {
		return resp, err
	}
	_ = redis.RedisClient.Set(s.ctx, cachedKey, ss, time.Minute*15)
	return resp, err
}
