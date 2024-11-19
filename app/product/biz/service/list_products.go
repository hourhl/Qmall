package service

import (
	"context"
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/product/biz/model"
	product "github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.

	categorytQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categorytQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductsResp{}

	for _, categoryList := range c {
		for _, p := range categoryList.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
			})
		}
	}

	return resp, nil
}
