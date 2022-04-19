package logic

import (
	"context"
	"go.uber.org/zap"
	"github.com/SunMaybo/zero/common/zlog"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
)

type SayStream1Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger *zap.SugaredLogger
}

func NewSayStream1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *SayStream1Logic {
	return &SayStream1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: zlog.WithContext(ctx),
	}
}
func (l *SayStream1Logic) SayStream1(stream greeter.GreeterService_SayStream1Server) error {
	return nil
}
