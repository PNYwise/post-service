package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/PNYwise/post-service/internal"
	"github.com/PNYwise/post-service/internal/config"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"
)

type ExtConf struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
}
type App struct {
	Port int `json:"port"`
}
type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Post     int    `json:"port"`
}

func main() {
	// Set time.Local to time.UTC
	time.Local = time.UTC

	// Initialize gRPC server
	srv := grpc.NewServer()

	// Load configuration
	conf := config.New()
	// Dial the gRPC server
	grpcConn, err := grpc.Dial(
		conf.GetString("config.host")+":"+conf.GetString("config.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Config Service gRPC server: %v", err)
	}
	defer grpcConn.Close()
	log.Println("Connected to Config Service gRPC server")

	// Create a gRPC client
	client := social_media_proto.NewConfigClient(grpcConn)

	// Create metadata
	md := metadata.New(map[string]string{
		"id":    conf.GetString("id"),
		"token": conf.GetString("token"),
	})

	// Add metadata to the context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Call the Get method on the server
	response, err := client.Get(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}

	// Parse the response
	extConf := &ExtConf{}
	if stringVal, ok := response.Kind.(*structpb.Value_StringValue); ok {
		err := json.Unmarshal([]byte(stringVal.StringValue), &extConf)
		if err != nil {
			log.Fatalf("Error unmarshaling configuration: %v", err)
		}
	}

	// Initialize gRPC server based on retrieved configuration
	internal.InitGrpc(srv)

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
