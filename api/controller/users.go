package controller

import (
	"github.com/gin-gonic/gin"
)

// GET /users
func (x *xController) GetUsers(ctx *gin.Context) {
	users, err := x.uu.GetMany(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	usersResp := toUserResponseSlice(users)
	ctx.JSON(200, usersResp)
}

// POST /users
func (x *xController) PostUsers(ctx *gin.Context) {
	var userReq UserCreate
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	newUser, err := x.uu.Create(ctx, userReq.Username, userReq.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	userResp := toUserResponse(*newUser)
	ctx.JSON(201, userResp)
}
