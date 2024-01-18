package repository

import (
	"github.com/PNYwise/post-service/internal/domain"
	"github.com/jackc/pgx/v5"
)

type postRepository struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.IPostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) Create(post *domain.Post) (*domain.Post, error) {
	return &domain.Post{
		Location: &domain.Location{},
	}, nil
}
func (p *postRepository) ReadAllByUserId(user_uuid string) (*[]domain.Post, error) {
	return nil, nil
}
func (p *postRepository) Delete(uuid string) error {
	return nil
}
