package dao

import (
	"utes-x-api/model"
	sqlboiler "utes-x-api/sqlboiler/entity"
)

func toPostModel(postDto sqlboiler.Post) (*model.Post, error) {
	post := model.Post{
		ID:        model.MustParseID(postDto.ID),
		UserID:    model.MustParseID(postDto.UserID),
		Content:   postDto.Content,
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

func toReplyModel(replyDto sqlboiler.Reply) (*model.Post, error) {
	reply := model.Post{
		ID:        model.MustParseID(replyDto.ID),
		UserID:    model.MustParseID(replyDto.UserID),
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
	Replies := make([]model.Post, len(repliesDto))
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
