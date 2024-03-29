package main

import (
	"flag"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/config"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/server"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
	"github.com/SunMaybo/zero/common/zcfg"
	"github.com/SunMaybo/zero/common/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

var cfgPath = flag.String("cfg", "/Users/fico/go/src/metadata/services/asset_platform/greeter/rpc/etc/config.yaml", "config file path")

func init() {
	flag.Parse()
}

func main() {
	cfg := config.Config{}
	zcfg.LoadConfig(*cfgPath, &cfg)
	//jwtInterceptor := grpc.ChainUnaryInterceptor(
	//	interceptor.UnaryJWTServerInterceptor("secret", nil),
	//))
	s := zrpc.NewServer(cfg.Zero, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    10 * time.Second,
		Timeout: 1 * time.Second,
	}), grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: true,
		MinTime:             10 * time.Second,
	}))
	defer s.Stop()
	s.RegisterServer(func(s *grpc.Server) error {
		serviceContext := svc.NewServiceContext(cfg)
		greeter.RegisterGreeterServiceServer(s, server.NewServer(serviceContext))
		return nil
	}).Start()
}
