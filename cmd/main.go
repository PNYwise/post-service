package main

import (
	"context"
	"time"

	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type postHandler struct {
	social_media_proto.UnimplementedPostServer
}

func NewPostHandler() *postHandler {
	return &postHandler{}
}

func (p *postHandler) Create(context.Context, *social_media_proto.PostDetail) (*social_media_proto.PostDetail, error) {
	data := &social_media_proto.PostDetail{}
	return data, nil
}
func (p *postHandler) ReadAllByUserId(context.Context, *social_media_proto.Uuid) (*social_media_proto.PostList, error) {
	data := &social_media_proto.PostList{}
	return data, nil
}
func (p *postHandler) Delete(context.Context, *social_media_proto.Uuid) (*empty.Empty, error) {
	return nil, nil
}

func main() {

	time.Local = time.UTC
	srv := grpc.NewServer()
	postHandlers := NewPostHandler()
	social_media_proto.RegisterPostServer(srv, postHandlers)

}
