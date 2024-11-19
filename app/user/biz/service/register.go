package service

import (
	"context"
	"errors"
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/biz/model"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 1. check whether email or password is empty or not
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("Email or password is empty")
	}

	// 2. check : password ?= password_confirm
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password not match")
	}

	// 3.generate passwordHash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	}

	// 4. insert user data
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
