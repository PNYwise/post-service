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
	err := p.postRepository.Create(post)
	return post, err
}
func (p *postService) ReadAllByUserId(userUuid string) (*[]domain.Post, error) {
	return p.postRepository.ReadAllByUserId(userUuid)
}
func (p *postService) Delete(uuid string) error {
	return nil
}
