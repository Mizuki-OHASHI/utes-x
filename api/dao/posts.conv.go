package dao

import (
	"utes-x-api/model"
	sqlboiler "utes-x-api/sqlboiler/entity"
)

func toPostModel(postDto sqlboiler.Post) (*model.Post, error) {
	likesDto := postDto.R.GetLikes()
	likes, err := toPostLikeModelSlice(likesDto)
	if err != nil {
		return nil, err
	}
	post := model.Post{
		ID:        model.MustParseID(postDto.ID),
		UserID:    model.MustParseID(postDto.UserID),
		Content:   postDto.Content,
		Likes:     likes,
		CreatedAt: postDto.CreatedAt,
		UpdatedAt: postDto.UpdatedAt.Ptr(),
	}
	return &post, nil
}

func toPostModelSlice(postsDto sqlboiler.PostSlice) ([]model.Post, error) {
	posts := make([]model.Post, len(postsDto))
	for i, postDto := range postsDto {
		post, err := toPostModel(*postDto)
		if err != nil {
			return nil, err
		}
		posts[i] = *post
	}
	return posts, nil
}

func toReplyModel(replyDto sqlboiler.Reply) (*model.Reply, error) {
	reply := model.Reply{
		ID:        model.MustParseID(replyDto.ID),
		UserID:    model.MustParseID(replyDto.UserID),
		PostID:    model.MustParseID(replyDto.PostID),
		Content:   replyDto.Content,
		CreatedAt: replyDto.CreatedAt,
		UpdatedAt: replyDto.UpdatedAt.Ptr(),
	}
	return &reply, nil
}

func toPostWithRepliesModel(postDto sqlboiler.Post) (*model.PostWithReplies, error) {
	post, err := toPostModel(postDto)
	if err != nil {
		return nil, err
	}
	repliesDto := postDto.R.GetReplies()
	Replies := make([]model.Reply, len(repliesDto))
	for i, replyDto := range repliesDto {
		reply, err := toReplyModel(*replyDto)
		if err != nil {
			return nil, err
		}
		Replies[i] = *reply
	}
	return &model.PostWithReplies{
		Post:    *post,
		Replies: Replies,
	}, nil
}

func toPostLikeModel(likeDto sqlboiler.Like) (*model.PostLike, error) {
	like := model.PostLike{
		ID:        model.MustParseID(likeDto.ID),
		PostID:    model.MustParseID(likeDto.PostID),
		UserID:    model.MustParseID(likeDto.UserID),
		CreatedAt: likeDto.CreatedAt,
		UpdatedAt: likeDto.UpdatedAt.Ptr(),
	}
	if likeDto.R.GetUser() != nil {
		user, err := toUserModel(*likeDto.R.GetUser())
		if err != nil {
			return nil, err
		}
		like.User = user
	}
	return &like, nil
}

func toPostLikeModelSlice(likesDto sqlboiler.LikeSlice) ([]model.PostLike, error) {
	likes := make([]model.PostLike, len(likesDto))
	for i, likeDto := range likesDto {
		like, err := toPostLikeModel(*likeDto)
		if err != nil {
			return nil, err
		}
		likes[i] = *like
	}
	return likes, nil
}
