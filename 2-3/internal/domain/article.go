package domain

import "time"

type Title string

type Content string

type Author string

type Article struct {
	ID        uint
	Title     Title
	Content   Content
	Author    Author
	CreatedAt time.Time
	UpdatedAt time.Time
}
