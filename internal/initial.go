package internal

import (
	"github.com/PNYwise/post-service/internal/handler"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server) {
	configGrpcInit(srv)
}

func configGrpcInit(srv *grpc.Server) {
	postHandlers := handler.NewPostHandler()
	social_media_proto.RegisterPostServer(srv, postHandlers)
}
