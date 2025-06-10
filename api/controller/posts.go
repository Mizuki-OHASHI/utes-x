package controller

import "github.com/gin-gonic/gin"

type postController struct {
	ServerInterface
}

func NewPostController() ServerInterface {
	return &postController{}
}

// GET /users/{user_id}/posts
func (p *postController) GetUsersUserIdPosts(ctx *gin.Context, userId string) {}

// POST /posts
func (p *postController) PostPosts(ctx *gin.Context) {}

// GET /posts/{post_id}
func (p *postController) GetPostsPostId(ctx *gin.Context, postId string) {}

// POST /reposts
func (p *postController) PostReposts(ctx *gin.Context) {}
