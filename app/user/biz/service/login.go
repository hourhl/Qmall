package service

import (
	"context"
	"errors"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/biz/model"
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
	// TODO
	token := "token"
	resp = &user.LoginResp{
		Token: token,
	}
	return resp, nil
}
