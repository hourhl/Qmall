package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hourhl/Qmall/app/auth/biz/model"
	auth "github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	"os"
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
	fmt.Printf("tokenString : %s", tokenString)
	// unit test
	//err = godotenv.Load("../../.env")

	if err != nil {
		kerrors.NewGRPCBizStatusError(1001, "Get env error")
		return nil, err
	}
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	fmt.Printf("mysignkey : %v\n", mySigningKey)

	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Print("ParseWithClaims failed\n")
		kerrors.NewGRPCBizStatusError(1002, "ParseWithClaims failed")
		resp = &auth.VerifyResp{Res: false}
		return resp, err
	} else if claim, ok := token.Claims.(*model.Claim); ok {
		id := claim.UserId
		if id != req.UserId {
			fmt.Print("Id in token and userId is not match\n")
			resp = &auth.VerifyResp{Res: false}
			return resp, nil
		}

	} else {
		kerrors.NewGRPCBizStatusError(1003, "unknown claims type, cannot proceed")
		resp = &auth.VerifyResp{Res: false}
		return resp, nil
	}

	fmt.Printf("Congratulation!!!\n")
	resp = &auth.VerifyResp{Res: true}
	return resp, nil
}
