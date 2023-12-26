package domain

type IPostRepository interface {
	Create(post *Post) (*Post, error)
	ReadAllByUserId(user_uuid string) (*[]Post, error)
	Delete(uuid string)
}
