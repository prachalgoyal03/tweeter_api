package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
//	"github.com/go-ozzo/ozzo-validation/v4/is"
)


type GetTweetsRequest struct {
	PaginationRequest
	ID  int `schema:"ID"`
	UserId     int `schema:"user_id"`
	ParentTweet int `schema:"parent_tweet"`
}

type CreateTweetRequest struct {
	UserId int `json:"user_id,omitempty"`
	Content  string `json:"content,omitempty"`
	ParentTweet int `json:"parent_tweet,omitempty"`
}

func (r CreateTweetRequest) ValidateCreateTweetRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserId, validation.Required),
		validation.Field(&r.Content, validation.Required),
	)
}

