package service

import (
	"context"
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
	product "github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	err := godotenv.Load("../../.env")
	mysql.Init()
	// init req and assert value

	req := &product.SearchProductsReq{
		Query: "shirt",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASS

}
