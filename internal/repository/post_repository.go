package repository

import (
	"context"
	"encoding/json"
	"log"

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
		log.Fatalf("error unmarshaling location: %v", err)
		return err
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
		log.Fatalf("err: %v", err)
		return err
	}
	err = json.Unmarshal(locationJSON, &post.Location)
	if err != nil {
		log.Fatalf("error unmarshaling location: %v", err)
		return err
	}

	return nil
}
func (p *postRepository) ReadAllByUserId(userUuid string) (*[]domain.Post, error) {
	query := `SELECT uuid, user_uuid, caption,image_url,location FROM posts p WHERE p.user_uuid = $1`

	rows, err := p.db.Query(context.Background(), query, &userUuid)
	if err != nil {
		log.Fatal("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		var post domain.Post
		var locationJSON []byte
		err := rows.Scan(&post.Uuid, &post.UserUuid, &post.Caption, &post.ImageUrl, &locationJSON)
		if err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, err
		}
		err = json.Unmarshal(locationJSON, &post.Location)
		if err != nil {
			log.Fatalf("error unmarshaling location: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		return nil, err
	}

	return &posts, nil
}
func (p *postRepository) Delete(uuid string) error {
	return nil
}
