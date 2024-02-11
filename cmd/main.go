package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"github.com/PNYwise/post-service/internal"
	"github.com/PNYwise/post-service/internal/config"
	"github.com/PNYwise/post-service/internal/domain"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	// Set time.Local to time.UTC
	time.Local = time.UTC

	// Initialize gRPC server
	srv := grpc.NewServer()

	// Load configuration
	conf := config.New()

	// Dial the gRPC server
	grpcConn, err := grpc.Dial(
		conf.GetString("config-service.host")+":"+conf.GetString("config-service.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Config Service gRPC server: %v", err)
	}
	log.Println("Connected to Config Service gRPC server")

	// Create a gRPC client
	client := social_media_proto.NewConfigClient(grpcConn)
	// Create metadata

	// Add metadata to the context
	ctx := createMetadataContext(conf)

	// Call the Get method on the server
	response, err := client.Get(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}
	grpcConn.Close()

	// Parse the response
	extConf, err := parseConfigResponse(response)
	if err != nil {
		log.Fatalf("Error unmarshaling configuration: %v", err)
	}

	// Initialize the db
	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		extConf.Database.Username,
		extConf.Database.Password,
		extConf.Database.Host,
		extConf.Database.Port,
		extConf.Database.Name,
	)
	connConfig, err := pgx.ParseConfig(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database")

	// Kafka broker address
	brokerList := []string{"127.0.0.1:9092"}

	// Initialize Kafka producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatal("error to creates a new sync producer")
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal("error closing Kafka producer:", err)
		}
	}()

	// Initialize gRPC server based on retrieved configuration
	internal.InitGrpc(srv, extConf, db, producer)

	// Start server
	serverPort := strconv.Itoa(extConf.App.Port)
	l, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("Could not listen to %s: %v", ":"+serverPort, err)
	}
	defer l.Close()

	log.Println("Server started at", ":"+serverPort)
	if err := srv.Serve(l); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func createMetadataContext(conf *viper.Viper) context.Context {
	// Add metadata to the context
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"id":    conf.GetString("id"),
		"token": conf.GetString("token"),
	}))
}

func parseConfigResponse(response *structpb.Value) (*domain.ExtConf, error) {
	extConf := &domain.ExtConf{}
	if stringVal, ok := response.Kind.(*structpb.Value_StringValue); ok {
		err := json.Unmarshal([]byte(stringVal.StringValue), extConf)
		return extConf, err
	}
	return nil, nil
}
