package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "github.com/hourhl/Qmall/app/frontend/hertz_gen/frontend/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "blackcat-1", "Price": 100, "Picture": "/static/image/blackcat.png"},
		{"Name": "threebody-1", "Price": 110, "Picture": "/static/image/threebody.png"},
		{"Name": "blackcat-2", "Price": 120, "Picture": "/static/image/blackcat.png"},
		{"Name": "threebody-2", "Price": 130, "Picture": "/static/image/threebody.png"},
		{"Name": "blackcat-3", "Price": 140, "Picture": "/static/image/blackcat.png"},
		{"Name": "threebody-3", "Price": 150, "Picture": "/static/image/threebody.png"},
	}

	resp["Title"] = "Hot Sale"
	resp["Items"] = items

	return resp, nil
}
