package controller

import "github.com/gin-gonic/gin"

type userController struct {
	ServerInterface
}

func NewUserController() ServerInterface {
	return &userController{}
}

// GET /users
func (u *userController) GetUsers(ctx *gin.Context) {}

// POST /users
func (u *userController) PostUsers(ctx *gin.Context) {}
