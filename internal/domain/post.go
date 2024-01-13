package domain

type Location struct {
	Lat float64
	Lng float64
}

type Post struct {
	Uuid     string
	UserUuid string
	Caption  string
	ImageUrl string
	Location *Location
}

type PostRequest struct {
	UserUuid string
	Caption  string
	ImageUrl string
	Location *Location
}

type IPostService interface {
	Create(post *PostRequest) (*Post, error)
	ReadAllByUserId(uuid string) (*Post, error)
	Delete(uuid string) error
}

type IPostRepository interface {
	Create(post *Post) (*Post, error)
	ReadAllByUserId(user_uuid string) (*[]Post, error)
	Delete(uuid string) error
}
