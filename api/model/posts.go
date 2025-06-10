package model

import "time"

type Post struct {
	ID        ID
	UserID    ID
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type PostWithReposts struct {
	Post
	Reposts []Post
}
