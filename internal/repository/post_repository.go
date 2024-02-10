package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
		log.Fatalf("error marshaling location: %v", err)
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
	query := `
	SELECT
	 uuid, user_uuid, caption,image_url,location 
	FROM posts p 
	WHERE p.user_uuid = $1 and deleted_at is null`

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

// Exist implements domain.IPostRepository.
func (p *postRepository) Exist(uuid string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM posts WHERE uuid = $1)"
	var exist bool
	row, err := p.db.Query(context.Background(), query, uuid)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return false, err
	}
	for row.Next() {
		if err := row.Scan(&exist); err != nil {
			log.Fatalf("Error Scaning query: %v", err)
			return false, err
		}
	}
	return exist, nil
}

func (p *postRepository) Delete(uuid string) error {
	query := `UPDATE posts SET deleted_at = $1 WHERE uuid = $2 and deleted_at is null`
	result, err := p.db.Exec(
		context.Background(),
		query,
		time.Now(),
		&uuid,
	)
	if err != nil {
		log.Fatalf("error executing query: %v", err)
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post with UUID %s not found", uuid)
	}

	return nil
}
