package models

type Coin struct {
	Country string
	Region  string
	Year    int
}

type News struct {
	Id         int
	CreatedAt  string
	Header     string
	Body       string
	PictureURL string
}
