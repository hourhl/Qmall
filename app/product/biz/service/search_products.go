package service

import (
	"context"
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/product/biz/model"
	product "github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var result []*product.Product
	for _, val := range products {
		result = append(result, &product.Product{
			Id:          uint32(val.ID),
			Name:        val.Name,
			Description: val.Description,
			Picture:     val.Picture,
			Price:       val.Price,
		})
	}
	return &product.SearchProductsResp{Results: result}, err
}
