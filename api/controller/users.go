package controller

import (
	"utes-x-api/usecase"

	"github.com/gin-gonic/gin"
)

type userController struct {
	ServerInterface
	uu usecase.User
}

func NewUserController(uu usecase.User) ServerInterface {
	return &userController{uu: uu}
}

// GET /users
func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.uu.GetMany(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	usersResp := toUserResponseSlice(users)
	ctx.JSON(200, usersResp)
}

// POST /users
func (u *userController) PostUsers(ctx *gin.Context) {
	var userReq UserCreate
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	newUser, err := u.uu.Create(ctx, userReq.Username, userReq.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	userResp := toUserResponse(*newUser)
	ctx.JSON(201, userResp)
}
