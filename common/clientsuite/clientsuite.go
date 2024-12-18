package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuit struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuit) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
	}

	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
