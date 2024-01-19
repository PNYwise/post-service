package domain

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Post struct {
	Uuid     string    `json:"uuid"`
	UserUuid string    `json:"user_uuid"`
	Caption  string    `json:"caption"`
	ImageUrl string    `json:"image_url"`
	Location *Location `json:"location"`
}

type PostRequest struct {
	UserUuid string    `json:"user_uuid"`
	Caption  string    `json:"caption"`
	ImageUrl string    `json:"image_url"`
	Location *Location `json:"location"`
}

type IPostService interface {
	Create(post *PostRequest) (*Post, error)
	ReadAllByUserId(uuid string) (*[]Post, error)
	Delete(uuid string) error
}

type IPostRepository interface {
	Create(post *Post) error
	ReadAllByUserId(userUuid string) (*[]Post, error)
	Delete(uuid string) error
}