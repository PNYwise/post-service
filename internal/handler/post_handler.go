package handler

import (
	"context"

	"github.com/PNYwise/post-service/internal/domain"
	social_media_proto "github.com/PNYwise/post-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type postHandler struct {
	social_media_proto.UnimplementedPostServer
	extConf     *domain.ExtConf
	postService domain.IPostService
}

func NewPostHandler(extConf *domain.ExtConf, postService domain.IPostService) *postHandler {
	return &postHandler{
		extConf:     extConf,
		postService: postService,
	}
}

func (p *postHandler) Create(_ context.Context, request *social_media_proto.PostDetail) (*social_media_proto.PostDetail, error) {
	// request
	postRequest := &domain.PostRequest{
		UserUuid: request.GetUserUuid(),
		Caption:  request.GetCaption(),
		ImageUrl: request.GetImageUrl(),
		Location: &domain.Location{
			Lat: request.Location.GetLat(),
			Lng: request.Location.GetLng(),
		},
	}

	// exec
	postResponse, err := p.postService.Create(postRequest)
	if err != nil {
		return nil, err
	}

	// response
	return &social_media_proto.PostDetail{
		UserUuid: postResponse.UserUuid,
		Caption:  postResponse.Caption,
		ImageUrl: postResponse.ImageUrl,
		Location: &social_media_proto.Location{
			Lat: postResponse.Location.Lat,
			Lng: postResponse.Location.Lng,
		},
	}, nil
}
func (p *postHandler) ReadAllByUserId(context.Context, *social_media_proto.Uuid) (*social_media_proto.PostList, error) {
	data := &social_media_proto.PostList{}
	return data, nil
}
func (p *postHandler) Delete(context.Context, *social_media_proto.Uuid) (*empty.Empty, error) {
	return nil, nil
}
