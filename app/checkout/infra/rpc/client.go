package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/hourhl/Qmall/app/checkout/conf"
	"github.com/hourhl/Qmall/common/clientsuite"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	PaymentClient paymentservice.Client
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	once.Do(func() {
		InitPaymentClient()
		InitProductClient()
		InitCartClient()
		InitOrderClient()
	})
}

func InitPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func InitCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func InitProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func InitOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
