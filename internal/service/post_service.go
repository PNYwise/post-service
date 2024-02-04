package service

import (
	"errors"

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

// Exist implements domain.IPostService.
func (p *postService) Exist(uuid string) (bool, error) {
	exist, err := p.postRepository.Exist(uuid)
	if err != nil {
		return false, errors.New("internal server error")
	}
	return exist, nil
}

func (p *postService) Delete(uuid string) error {
	return p.postRepository.Delete(uuid)
}
