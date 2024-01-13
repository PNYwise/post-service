package repository

import "github.com/PNYwise/post-service/internal/domain"

type postRepository struct {
}

func NewPostRepository() domain.IPostRepository {
	return &postRepository{}
}

func (p *postRepository) Create(post *domain.Post) (*domain.Post, error) {
	return nil, nil
}
func (p *postRepository) ReadAllByUserId(user_uuid string) (*[]domain.Post, error) {
	return nil, nil
}
func (p *postRepository) Delete(uuid string) error {
	return nil
}
