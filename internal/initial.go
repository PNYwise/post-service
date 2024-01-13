package internal

import (
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/PNYwise/post-service/internal/handler"
	"github.com/PNYwise/post-service/internal/repository"
	"github.com/PNYwise/post-service/internal/service"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, extConf *domain.ExtConf) {
	postRepostory := repository.NewPostRepository()
	postService := service.NewPostService(postRepostory)
	postHandlers := handler.NewPostHandler(extConf, postService)
	social_media_proto.RegisterPostServer(srv, postHandlers)
}
