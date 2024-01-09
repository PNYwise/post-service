package main

import (
	"context"
	"log"
	"time"

	"github.com/PNYwise/post-service/internal"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	time.Local = time.UTC
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := social_media_proto.NewConfigClient(conn)

	md := metadata.New(map[string]string{
		"id":    "post-service",
		"token": "a8b90bf5e550b538b82758b970750341",
	})
	// Add metadata to the context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Call the Get method on the server
	response, err := client.Get(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error calling Get: %v", err)
	}
	if stringVal, ok := response.Kind.(*structpb.Value_StringValue); ok {
		val := stringVal
		log.Printf(val.StringValue)
	}
	srv := grpc.NewServer()
	internal.InitGrpc(srv)

}
