package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hourhl/Qmall/app/auth/biz/model"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"github.com/joho/godotenv"
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

	err = godotenv.Load()
	if err != nil {
		kerrors.NewGRPCBizStatusError(1001, "Get env error")
		return nil, err
	}
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	ss, err := token.SignedString(mySigningKey)

	resp = &auth.DeliveryResp{Token: ss}
	return resp, err
}
