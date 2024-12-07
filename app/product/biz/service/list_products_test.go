package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
	product "github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("load env error %s \n", err.Error())
	}
	mysql.Init()
	// init req and assert value

	req := &product.ListProductsReq{
		Page:         2,
		PageSize:     1,
		CategoryName: "T-Shirt",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// status : PASSS

}
