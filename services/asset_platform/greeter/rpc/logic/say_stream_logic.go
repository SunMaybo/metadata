package logic

import (
	"context"
	"go.uber.org/zap"
	"github.com/SunMaybo/zero/common/zlog"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
)

type SayStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger *zap.SugaredLogger
}

func NewSayStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayStreamLogic {
	return &SayStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: zlog.WithContext(ctx),
	}
}
func (l *SayStreamLogic) SayStream(stream greeter.GreeterService_SayStreamServer) error {
	return nil
}
