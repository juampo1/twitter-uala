package dto

type UserFollowRequest struct {
	FollowedID string `json:"userId" validate:"required"`
}
