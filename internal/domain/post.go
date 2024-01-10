package domain

import social_media_proto "github.com/PNYwise/post-service/proto"

type Location struct {
	Lat float64
	Lng float64
}

type Post struct {
	Uuid     string
	UserUuid string
	Caption  string
	ImageUrl string
	Location Location
}

type IPostService interface {
	Create(post *social_media_proto.PostDetail) (*social_media_proto.PostDetail, error)
	ReadAllByUserId(user_uuid *social_media_proto.Uuid) (*social_media_proto.PostList, error)
	Delete(uuid *social_media_proto.Uuid) error
}

type IPostRepository interface {
	Create(post *Post) (*Post, error)
	ReadAllByUserId(user_uuid string) (*[]Post, error)
	Delete(uuid string) error
}
