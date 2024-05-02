package imagerecognition

import (
	"net/http"
	"net/url"
)

type Service interface {
	Dish(in DishRequest) (DishResponse, error)
}

type service struct {
	accessToken url.Values
	client      *http.Client
}

func NewImageRecognition(accessToken url.Values, client *http.Client) Service {
	return &service{
		accessToken: accessToken,
		client:      client,
	}
}

var _ Service = (*service)(nil)
