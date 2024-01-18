package repository

import (
	"context"
	"encoding/json"
	"fmt"

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

func (p *postRepository) Create(post *domain.Post) error {

	locationJSON, err := json.Marshal(post.Location)
	if err != nil {
		return fmt.Errorf("error marshaling location: %v", err)
	}

	query := `
	INSERT INTO posts (user_uuid, caption, image_url, location)
	VALUES ($1, $2, $3, $4)
	RETURNING uuid, user_uuid, caption, image_url, location
`
	err = p.db.QueryRow(
		context.Background(),
		query,
		post.UserUuid, post.Caption, post.ImageUrl, locationJSON,
	).Scan(&post.Uuid, &post.UserUuid, &post.Caption, &post.ImageUrl, &locationJSON)
	if err != nil {
		return err
	}
	err = json.Unmarshal(locationJSON, &post.Location)
	if err != nil {
		return fmt.Errorf("error unmarshaling location: %v", err)
	}

	return nil
}
func (p *postRepository) ReadAllByUserId(user_uuid string) (*[]domain.Post, error) {
	return nil, nil
}
func (p *postRepository) Delete(uuid string) error {
	return nil
}
