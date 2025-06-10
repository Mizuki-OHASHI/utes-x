package controller

import "github.com/gin-gonic/gin"

type repostController struct {
	ServerInterface
}

func newRepostController() ServerInterface {
	return &repostController{}
}

// POST /reposts
func (r *repostController) PostReposts(ctx *gin.Context) {}
