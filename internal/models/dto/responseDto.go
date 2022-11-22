package dto

type NewsResponseBody struct {
	Header     string `json:"header"`
	Body       string `json:"body"`
	PictureURL string `json:"pictureURL"`
}

type NewsResponseParams struct {
	NewsID uint `json:"newsID"`
}
