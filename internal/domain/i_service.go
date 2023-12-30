package domain

import social_media_proto "github.com/PNYwise/post-service/proto"

type IPostService interface {
	Create(post *social_media_proto.PostDetail) (*social_media_proto.PostDetail, error)
	ReadAllByUserId(user_uuid *social_media_proto.Uuid) (*social_media_proto.PostList, error)
	Delete(uuid *social_media_proto.Uuid) error
}
