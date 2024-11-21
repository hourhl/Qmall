package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/hourhl/Qmall/app/checkout/infra/rpc"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart"
	checkout "github.com/hourhl/Qmall/rpc_gen/kitex_gen/checkout"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(6005001, err.Error())
	}

	if cartResult == nil || cartResult.Cart == nil {
		return nil, kerrors.NewGRPCBizStatusError(6005002, "cart is empty")
	}

	var total float32
	// TODO
	// 应该在for循环外面进行rpc调用，否则性能不太好
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resulterr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resulterr != nil {
			return nil, kerrors.NewGRPCBizStatusError(6005003, "cannot find product")
		}
		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price
		total += p * float32(cartItem.Quantity)
	}

	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()

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

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})

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
