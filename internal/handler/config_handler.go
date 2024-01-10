package handler

import (
	"context"
	"fmt"

	"github.com/PNYwise/post-service/internal/domain"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type postHandler struct {
	social_media_proto.UnimplementedPostServer
	extConf *domain.ExtConf
}

func NewPostHandler(extConf *domain.ExtConf) *postHandler {
	return &postHandler{
		extConf: extConf,
	}
}

func (p *postHandler) Create(context.Context, *social_media_proto.PostDetail) (*social_media_proto.PostDetail, error) {
	data := &social_media_proto.PostDetail{}
	fmt.Print(p.extConf)
	return data, nil
}
func (p *postHandler) ReadAllByUserId(context.Context, *social_media_proto.Uuid) (*social_media_proto.PostList, error) {
	data := &social_media_proto.PostList{}
	return data, nil
}
func (p *postHandler) Delete(context.Context, *social_media_proto.Uuid) (*empty.Empty, error) {
	return nil, nil
}
