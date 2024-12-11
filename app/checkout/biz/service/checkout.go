package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hourhl/Qmall/app/checkout/infra/rpc"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	checkout "github.com/hourhl/Qmall/rpc_gen/kitex_gen/checkout"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/order"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId, Token: req.Token})
	if err != nil {
		fmt.Printf("get cart failed, err:%v\n", err)
		return nil, kerrors.NewGRPCBizStatusError(6005001, err.Error())
	}
	fmt.Printf("get cart succeed, cartResult : %v\n", cartResult)

	if cartResult == nil || cartResult.Cart == nil {
		fmt.Printf("cart is empty\n")
		return nil, kerrors.NewGRPCBizStatusError(6005002, "cart is empty")
	}

	var total float32
	var oi []*order.OrderItem
	// TODO
	// 应该在for循环外面进行rpc调用，否则性能不太好
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resulterr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resulterr != nil {
			fmt.Printf("get product failed, err:%v\n", resulterr)
			return nil, kerrors.NewGRPCBizStatusError(6005003, "cannot find product")
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	var orderId string
	zipInt, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			Street:   req.Address.StreetAddress,
			City:     req.Address.City,
			Province: req.Address.Province,
			Country:  req.Address.Country,
			ZipCode:  int32(zipInt),
		},
		OrderItems: oi,
	})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(6004001, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId, Token: req.Token})

	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{OrderId: orderId, TransactionId: paymentResult.TransactionId}
	return
}
