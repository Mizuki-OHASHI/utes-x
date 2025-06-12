package model

import "time"

type Post struct {
	ID        ID
	UserID    ID
	Content   string
	Likes     []PostLike
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Reply struct {
	ID        ID
	UserID    ID
	PostID    ID
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type PostWithReplies struct {
	Post
	Replies []Reply
}

type PostLike struct {
	ID        ID
	PostID    ID
	UserID    ID
	CreatedAt time.Time
	UpdatedAt *time.Time
	User      *User
}
