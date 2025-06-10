package controller

import "utes-x-api/usecase"

type xController struct {
	ServerInterface
	uu usecase.User
	up usecase.Post
}

func NewController(uu usecase.User, up usecase.Post) ServerInterface {
	return &xController{
		uu: uu,
		up: up,
	}
}
