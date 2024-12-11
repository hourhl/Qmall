package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/biz/model"
	"github.com/hourhl/Qmall/app/user/infra/rpc"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 1. check email or password
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Email or password is empty")
	}

	// 2. identification
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// 3. authorization
	if rpc.AuthClient == nil {
		println("authclient is nil")
		return nil, errors.New("authclient is nil")

	}
	token, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: int32(row.ID)})
	if err != nil {
		fmt.Printf("deliver token err: %v\n", err)
	}
	resp = &user.LoginResp{
		Token: token.Token,
	}

	// unit test
	//fmt.Printf("now test verify token\n")
	//verify, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{Token: token.Token})
	//if err != nil {
	//	fmt.Printf("verify token err: %v\n", err)
	//}
	//fmt.Printf("verify result is %v\n", verify.Res)
	return resp, nil

}
