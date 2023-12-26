package domain

type IPostService interface {
	Create(post *PostRequest) (*PostResponse, error)
	ReadAllByUserId(user_uuid string) (*[]PostResponse, error)
	Delete(uuid string) error
}
