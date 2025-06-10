package controller

import (
	"utes-x-api/model"

	"github.com/gin-gonic/gin"
)

// GET /users/{user_id}/posts
func (x *xController) GetUsersUserIdPosts(ctx *gin.Context, userId string) {
	userID, err := model.ParseID(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	posts, err := x.up.GetMany(ctx, userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch posts"})
		return
	}
	postsResp := toPostResponseSlice(posts)
	ctx.JSON(200, postsResp)
}

// POST /posts
func (x *xController) PostPosts(ctx *gin.Context) {
	var postReq PostCreate
	if err := ctx.ShouldBindJSON(&postReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	userID, err := model.ParseID(postReq.UserId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	newPost, err := x.up.Create(ctx, userID, postReq.Content)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}
	postResp := toPostResponse(*newPost)
	ctx.JSON(201, postResp)
}

// GET /posts/{post_id}
func (x *xController) GetPostsPostId(ctx *gin.Context, postId string) {
	postID, err := model.ParseID(postId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}
	postWithReplies, err := x.up.GetWithReplies(ctx, postID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch post with replies"})
		return
	}
	postWithRepliesResp := toPostWithRepliesResponse(*postWithReplies)
	ctx.JSON(200, postWithRepliesResp)
}

// POST /Replies
func (x *xController) PostReplies(ctx *gin.Context) {
	var replyReq ReplyCreate
	if err := ctx.ShouldBindJSON(&replyReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	replyTo, err := model.ParseID(replyReq.PostId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid reply to post ID"})
		return
	}
	userID, err := model.ParseID(replyReq.UserId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	newReply, err := x.up.CreateReply(ctx, replyTo, userID, replyReq.Content)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create reply"})
		return
	}
	replyResp := toReplyResponse(*newReply)
	ctx.JSON(201, replyResp)
}
