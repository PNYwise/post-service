package internal

import (
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/PNYwise/post-service/internal/handler"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, extConf *domain.ExtConf) {
	postHandlers := handler.NewPostHandler(extConf)
	social_media_proto.RegisterPostServer(srv, postHandlers)
}
