package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	PaymentClient paymentservice.Client
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
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
	var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	// unit test
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout"}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func InitCartClient() {
	var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	// unit test
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout"}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func InitProductClient() {
	var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	// unit test
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout"}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func InitOrderClient() {
	var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	// unit test
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout"}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
