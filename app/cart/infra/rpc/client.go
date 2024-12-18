package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hourhl/Qmall/app/cart/conf"
	cartutils "github.com/hourhl/Qmall/app/cart/utils"
	"github.com/hourhl/Qmall/common/clientsuite"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/user/userservice"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	once          sync.Once
	err           error

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	once.Do(func() {
		initProductClient()
		InitUserClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatal(err)
	}
	cartutils.MustHandlerError(err)
}

func InitUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		klog.Fatal(err)
	}
	cartutils.MustHandlerError(err)

}
