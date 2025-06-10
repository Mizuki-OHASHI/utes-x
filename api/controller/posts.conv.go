package controller

import (
	"utes-x-api/model"
)

func toPostResponse(post model.Post) Post {
	return Post{
		Id:        post.ID.String(),
		UserId:    post.UserID.String(),
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: nil,
	}
}

func toPostResponseSlice(posts []model.Post) []Post {
	res := make([]Post, len(posts))
	for i, post := range posts {
		res[i] = toPostResponse(post)
	}
	return res
}

func toReplyResponse(reply model.Post) Reply {
	return Reply{
		Id:        reply.ID.String(),
		UserId:    reply.UserID.String(),
		Content:   reply.Content,
		CreatedAt: reply.CreatedAt,
		UpdatedAt: nil,
	}
}

func toReplyResponseSlice(replies []model.Post) []Reply {
	res := make([]Reply, len(replies))
	for i, reply := range replies {
		res[i] = toReplyResponse(reply)
	}
	return res
}

func toPostWithRepliesResponse(post model.PostWithReplies) PostWithReplies {
	postResp := toPostResponse(post.Post)
	repliesResp := toReplyResponseSlice(post.Replies)
	return PostWithReplies{
		Post:    postResp,
		Replies: repliesResp,
	}
}
