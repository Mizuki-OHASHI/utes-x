package usecase

import (
	"context"
	"utes-x-api/dao"
	"utes-x-api/model"
)

type Post interface {
	GetMany(ctx context.Context, userID model.ID) ([]model.Post, error)
	Create(ctx context.Context, userID model.ID, content string) (*model.Post, error)
	GetWithReplies(ctx context.Context, postID model.ID) (*model.PostWithReplies, error)
	CreateReply(ctx context.Context, replyTo model.ID, userID model.ID, content string) (*model.Reply, error)
}

type postUsecase struct {
	pd dao.Post
}

func NewPostUsecase(pd dao.Post) Post {
	return &postUsecase{pd: pd}
}

func (p *postUsecase) GetMany(ctx context.Context, userID model.ID) ([]model.Post, error) {
	posts, err := p.pd.GetMany(ctx, dao.GetManyQuery{UserID: userID})
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postUsecase) Create(ctx context.Context, userID model.ID, content string) (*model.Post, error) {
	post := model.Post{
		ID:      model.NewID(),
		UserID:  userID,
		Content: content,
	}
	newPost, err := p.pd.Create(ctx, post)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (p *postUsecase) GetWithReplies(ctx context.Context, postID model.ID) (*model.PostWithReplies, error) {
	postWithReplies, err := p.pd.GetWithReplies(ctx, postID)
	if err != nil {
		return nil, err
	}
	return postWithReplies, nil
}

func (p *postUsecase) CreateReply(ctx context.Context, replyTo model.ID, userID model.ID, content string) (*model.Reply, error) {
	reply := model.Post{
		ID:      model.NewID(),
		UserID:  userID,
		Content: content,
	}
	newReply, err := p.pd.CreateReply(ctx, replyTo, userID, reply)
	if err != nil {
		return nil, err
	}
	return newReply, nil
}
