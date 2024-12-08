package rpc

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	cartutils "github.com/hourhl/Qmall/app/cart/utils"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
		InitUserClient()
	})
}

func initProductClient() {
	var opts []client.Option
	// dev
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")

	cartutils.MustHandlerError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		fmt.Printf("init product client error: %v\n", err)
	}
	cartutils.MustHandlerError(err)
}

func InitUserClient() {
	var opts []client.Option
	// dev
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		fmt.Printf("init consul resolver error: %v\n", err)
	}
	cartutils.MustHandlerError(err)

	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		fmt.Printf("init user client error: %v\n", err)
	}
	cartutils.MustHandlerError(err)

}
