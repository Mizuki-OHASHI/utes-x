package model

import "time"

type Post struct {
	ID        ID
	UserID    ID
	Content   string
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
