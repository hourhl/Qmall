package service

import (
	"context"
	"fmt"
	"github.com/hourhl/Qmall/app/product/biz/dal"
	product "github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	err := godotenv.Load("../../.env")
	dal.Init()
	// init req and assert value

	req := &product.GetProductReq{
		Id: 3,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	fmt.Printf("check the information : %v\n", resp.Product.Name == "Notebook")

	// status : PASS

}
