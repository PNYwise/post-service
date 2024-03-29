package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/PNYwise/post-service/internal"
	"github.com/PNYwise/post-service/internal/config"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Set time.Local to time.UTC
	time.Local = time.UTC

	// Initialize gRPC server
	srv := grpc.NewServer()

	// Load configuration
	conf := config.New()

	// Add metadata to the context
	ctx := createMetadataContext(conf)
	extConf := config.ConfigFromGrpcServer(ctx, conf)

	// Initialize the db
	db := config.DbConn(ctx, extConf)
	defer db.Close(ctx)

	// Initialize Kafka producer configuration
	producer := config.GetKafkaProducer(extConf)
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %v", err)
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
