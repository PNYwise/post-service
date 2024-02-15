package internal

import (
	"github.com/IBM/sarama"
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/PNYwise/post-service/internal/handler"
	"github.com/PNYwise/post-service/internal/repository"
	"github.com/PNYwise/post-service/internal/service"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, extConf *domain.ExtConf, db *pgx.Conn, producer sarama.SyncProducer) {
	kafkaPostRepository := repository.NewKafkaPostRepository(producer, extConf)
	postRepostory := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepostory, kafkaPostRepository)
	postHandlers := handler.NewPostHandler(extConf, postService)
	social_media_proto.RegisterPostServer(srv, postHandlers)
}
