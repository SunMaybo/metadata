package server

import (
	"context"
	"github.com/SunMaybo/metadata/proto/services/asset_platform/greeter"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/logic"
	"github.com/SunMaybo/metadata/services/asset_platform/greeter/rpc/svc"
)

type Server struct {
	svcCtx *svc.ServiceContext
    greeter.UnimplementedGreeterServiceServer
}

func NewServer(svcCtx *svc.ServiceContext) *Server {
	return &Server{
		svcCtx: svcCtx,
	}
}

func (s *Server) SayHelloWord(ctx context.Context, in *greeter.HelloRequest)  (*greeter.HelloReply, error){
	
	l := logic.NewSayHelloWordLogic(ctx, s.svcCtx)
	
	return l.SayHelloWord(in)
}

func (s *Server) SayStream(stream greeter.GreeterService_SayStreamServer) error{
	
	l := logic.NewSayStreamLogic(stream.Context(), s.svcCtx)
	
	return l.SayStream(stream)
}

func (s *Server) SayStream1(stream greeter.GreeterService_SayStream1Server) error{
	
	l := logic.NewSayStream1Logic(stream.Context(), s.svcCtx)
	
	return l.SayStream1(stream)
}

func (s *Server) SayStream2(in *greeter.HelloRequest, stream greeter.GreeterService_SayStream2Server) error{
	
	l := logic.NewSayStream2Logic(stream.Context(), s.svcCtx)
	
	return l.SayStream2(in,stream)
}


