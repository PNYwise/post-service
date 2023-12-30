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
	Location Location
}