package main

import (
	"github.com/hourhl/Qmall/app/checkout/biz/dal"
	"github.com/hourhl/Qmall/app/checkout/infra/rpc"
	"github.com/hourhl/Qmall/common/mtl"
	"github.com/hourhl/Qmall/common/serversuite"
	"github.com/joho/godotenv"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/hourhl/Qmall/app/checkout/conf"
	"github.com/hourhl/Qmall/rpc_gen/kitex_gen/checkout/checkoutservice"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	godotenv.Load(".env")
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	dal.Init()
	rpc.Init()
	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
