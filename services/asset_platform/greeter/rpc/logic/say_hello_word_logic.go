package logic

import (
	"context"
	"go.uber.org/zap"
	"github.com/SunMaybo/zero/common/zlog"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
)

type SayHelloWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger *zap.SugaredLogger
}

func NewSayHelloWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloWordLogic {
	return &SayHelloWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: zlog.WithContext(ctx),
	}
}
func (l *SayHelloWordLogic) SayHelloWord(in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{} ,nil
}
