package rpc

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/hourhl/Qmall/app/user/conf"
	"github.com/hourhl/Qmall/common/clientsuite"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth/authservice"
	"sync"
)

var (
	AuthClient authservice.Client
	once       sync.Once
	err        error

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	once.Do(func() {
		InitAuthClient()
	})
}

func InitAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuit{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		fmt.Sprintf("init authClient failed, err: %v\n", err)
		panic(err)
	}

}
