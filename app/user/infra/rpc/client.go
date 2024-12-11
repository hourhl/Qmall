package rpc

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/auth/authservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"

	"github.com/hourhl/Qmall/app/user/conf"
)

var (
	AuthClient authservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		InitAuthClient()
	})
}

func InitAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// unit test
	//r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		//panic(err)
		fmt.Printf("init authClient resolver failed, err: %v\n", err)
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	// dev
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)

	// unit test
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		fmt.Sprintf("init authClient failed, err: %v\n", err)
		panic(err)
	}

}
