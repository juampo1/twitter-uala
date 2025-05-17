package dto

type UserFollowRequest struct {
	UserToFollowID string `json:"userId" validate:"required"`
}
