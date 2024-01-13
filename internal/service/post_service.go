package service

import (
	"github.com/PNYwise/post-service/internal/domain"
)

type postService struct {
	postRepository domain.IPostRepository
}

func NewPostService(postRepository domain.IPostRepository) domain.IPostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (p *postService) Create(request *domain.PostRequest) (*domain.Post, error) {
	post := &domain.Post{
		UserUuid: request.UserUuid,
		Caption:  request.Caption,
		ImageUrl: request.ImageUrl,
		Location: request.Location,
	}
	data, err := p.postRepository.Create(post)
	return data, err
}
func (p *postService) ReadAllByUserId(uuid string) (*domain.Post, error) {
	return nil, nil
}
func (p *postService) Delete(uuid string) error {
	return nil
}
