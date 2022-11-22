package dto

type NewsRequestBody struct {
	Header     string `json:"header"`
	Body       string `json:"body"`
	PictureURL string `json:"pictureURL"`
}

type NewsRequestParams struct {
	NewsID uint
}
