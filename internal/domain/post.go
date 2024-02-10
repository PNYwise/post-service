package domain

type Location struct {
	Lat float64 `json:"lat" validate:"number"`
	Lng float64 `json:"lng" validate:"number"`
}

type Post struct {
	Uuid     string    `json:"uuid"`
	UserUuid string    `json:"user_uuid"`
	Caption  string    `json:"caption"`
	ImageUrl string    `json:"image_url"`
	Location *Location `json:"location"`
}

type PostRequest struct {
	UserUuid string    `json:"user_uuid" validate:"required,uuid4"`
	Caption  string    `json:"caption" validate:"-"`
	ImageUrl string    `json:"image_url"`
	Location *Location `json:"location"`
}

type IPostService interface {
	Create(post *PostRequest) (*Post, error)
	ReadAllByUserId(uuid string) (*[]Post, error)
	Exist(uuid string) (bool, error)
	Delete(uuid string) error
}

type IPostRepository interface {
	Create(post *Post) error
	ReadAllByUserId(userUuid string) (*[]Post, error)
	Exist(uuid string) (bool, error)
	Delete(uuid string) error
}

type KafkaPostRepository interface {
	PublishMessage(post *Post) error
}
