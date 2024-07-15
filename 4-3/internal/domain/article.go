package domain

import (
	"errors"
	"fmt"
	"time"
)

type ArticleID uint

type ArticleTitle string

const (
	MaxTitleLength  = 255
	MaxAuthorLength = 64
)

func NewTitle(title string) (ArticleTitle, error) {
	if len(title) > MaxTitleLength {
		return "", fmt.Errorf("title length must be less than %d", MaxTitleLength)
	}
	if title == "" {
		return ArticleTitle("No Title"), nil
	}
	return ArticleTitle(title), nil
}

type Content string

func NewContent(content string) (Content, error) {
	if content == "" {
		return "", errors.New("content cannot be empty")
	}
	return Content(content), nil
}

type Author string

func NewAuthor(author string) (Author, error) {
	if author == "" {
		return "", errors.New("author cannot be empty")
	}
	if len(author) > MaxAuthorLength {
		return "", fmt.Errorf("author length must be less than %d", MaxAuthorLength)
	}
	return Author(author), nil
}

type Article struct {
	ID        ArticleID
	Title     ArticleTitle
	Content   Content
	Author    Author
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Article) Update(title, content string) error {
	t, err := NewTitle(title)
	if err != nil {
		return err
	}
	c, err := NewContent(content)
	if err != nil {
		return err
	}
	a.Title = t
	a.Content = c
	// ここでtime.Now() してしまうと、テストがむずかしくなる
	a.UpdatedAt = time.Now()
	return nil
}

func newArticle(id ArticleID, title ArticleTitle, content Content, author Author, createdAt, updatedAt time.Time) *Article {
	return &Article{
		ID:        id,
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func CreateArticle(id uint, title, content, author string) (*Article, error) {
	i_d := ArticleID(id)
	t, err := NewTitle(title)
	if err != nil {
		return nil, err
	}
	c, err := NewContent(content)
	if err != nil {
		return nil, err
	}
	a, err := NewAuthor(author)
	if err != nil {
		return nil, err
	}
	// ここでtime.Now() してしまうと、テストがむずかしくなる
	now := time.Now()
	return newArticle(i_d, t, c, a, now, now), nil
}
