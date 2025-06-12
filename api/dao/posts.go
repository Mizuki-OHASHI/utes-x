package dao

import (
	"context"
	"database/sql"
	"utes-x-api/model"
	sqlboiler "utes-x-api/sqlboiler/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/xerrors"
)

type GetManyQuery struct {
	UserID model.ID
}

type Post interface {
	GetMany(ctx context.Context, query GetManyQuery) ([]model.Post, error)
	Create(ctx context.Context, post model.Post) (*model.Post, error)
	CreateReply(ctx context.Context, replyTo model.ID, userID model.ID, reply model.Post) (*model.Reply, error)
	GetWithReplies(ctx context.Context, postID model.ID) (*model.PostWithReplies, error)
	CreateLike(ctx context.Context, like model.PostLike) (*model.PostLike, error)
}

type postDao struct {
	db *sql.DB
}

func NewPostDao(db *sql.DB) Post {
	return &postDao{db: db}
}

func (p *postDao) GetMany(ctx context.Context, query GetManyQuery) ([]model.Post, error) {
	posts, err := sqlboiler.Posts(
		sqlboiler.PostWhere.UserID.EQ(query.UserID.String()),
		qm.Load(sqlboiler.PostRels.Likes),
		qm.Load(qm.Rels(sqlboiler.PostRels.Likes, sqlboiler.LikeRels.User)),
	).All(ctx, p.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to get posts: %w", err)
	}
	return toPostModelSlice(posts)
}

func (p *postDao) Create(ctx context.Context, post model.Post) (*model.Post, error) {
	postDto := sqlboiler.Post{
		ID:      post.ID.String(),
		UserID:  post.UserID.String(),
		Content: post.Content,
	}
	if err := postDto.Insert(ctx, p.db, boil.Infer()); err != nil {
		return nil, xerrors.Errorf("failed to create post: %w", err)
	}
	newPostDto, err := sqlboiler.Posts(
		sqlboiler.PostWhere.ID.EQ(post.ID.String()),
	).One(ctx, p.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to fetch created post: %w", err)
	}
	return toPostModel(*newPostDto)
}

func (p *postDao) CreateReply(ctx context.Context, replyTo model.ID, userID model.ID, reply model.Post) (*model.Reply, error) {
	replyDto := sqlboiler.Reply{
		ID:      reply.ID.String(),
		PostID:  replyTo.String(),
		UserID:  reply.UserID.String(),
		Content: reply.Content,
	}
	if err := replyDto.Insert(ctx, p.db, boil.Infer()); err != nil {
		return nil, xerrors.Errorf("failed to create reply: %w", err)
	}
	newReplyDto, err := sqlboiler.Replies(
		sqlboiler.ReplyWhere.ID.EQ(reply.ID.String()),
	).One(ctx, p.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to fetch created reply: %w", err)
	}
	return toReplyModel(*newReplyDto)
}

func (p *postDao) GetWithReplies(ctx context.Context, postID model.ID) (*model.PostWithReplies, error) {
	postDto, err := sqlboiler.Posts(
		sqlboiler.PostWhere.ID.EQ(postID.String()),
		qm.Load(sqlboiler.PostRels.Replies),
		qm.Load(sqlboiler.PostRels.Likes),
		qm.Load(qm.Rels(sqlboiler.PostRels.Likes, sqlboiler.LikeRels.User)),
	).One(ctx, p.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to get post with Replies: %w", err)
	}
	return toPostWithRepliesModel(*postDto)
}

func (p *postDao) CreateLike(ctx context.Context, like model.PostLike) (*model.PostLike, error) {
	likeDto := sqlboiler.Like{
		ID:     like.ID.String(),
		PostID: like.PostID.String(),
		UserID: like.UserID.String(),
	}
	if err := likeDto.Insert(ctx, p.db, boil.Infer()); err != nil {
		return nil, xerrors.Errorf("failed to create post like: %w", err)
	}
	newLikeDto, err := sqlboiler.Likes(
		sqlboiler.LikeWhere.ID.EQ(like.ID.String()),
	).One(ctx, p.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to fetch created post like: %w", err)
	}
	return toPostLikeModel(*newLikeDto)
}
