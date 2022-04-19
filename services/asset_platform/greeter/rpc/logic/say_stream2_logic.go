package logic

import (
	"context"
	"go.uber.org/zap"
	"github.com/SunMaybo/zero/common/zlog"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
)

type SayStream2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger *zap.SugaredLogger
}

func NewSayStream2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *SayStream2Logic {
	return &SayStream2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: zlog.WithContext(ctx),
	}
}
func (l *SayStream2Logic) SayStream2(in *greeter.HelloRequest, stream greeter.GreeterService_SayStream2Server) error {
	return nil
}
