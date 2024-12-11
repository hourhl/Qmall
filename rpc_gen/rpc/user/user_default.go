package user

import (
	"context"
	user "github.com/hourhl/Qmall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyUser(ctx context.Context, req *user.VerifyUserReq, callOptions ...callopt.Option) (resp *user.VerifyUserResp, err error) {
	resp, err = defaultClient.VerifyUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
