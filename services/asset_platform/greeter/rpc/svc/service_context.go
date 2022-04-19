package svc

import "github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(cfg config.Config) *ServiceContext {
	return &ServiceContext{
		Config: cfg,
	}
}
