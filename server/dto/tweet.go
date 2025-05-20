package dto

type UserTweetRequest struct {
	Content string `json:"content" validate:"required,min=1,max=280"`
}
