package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/u-masato/blogger/4-2/internal/domain"
)

type SQLArticleRepository struct {
	db *sql.DB
}

func NewSQLArticleRepository(db *sql.DB) *SQLArticleRepository {
	return &SQLArticleRepository{db: db}
}

func (r *SQLArticleRepository) Get(ctx context.Context, id domain.ArticleID) (*domain.Article, error) {
	var (
		aid                    uint
		title, content, author string
		created, updated       time.Time
	)

	row := r.db.QueryRowContext(ctx,
		`SELECT id, title, content, author, created, updated
		FROM articles WHERE id = ?`, id)
	err := row.Scan(&aid, &title, &content, &author, &created, &updated)
	if err != nil {
		return nil, err
	}
	article := &domain.Article{
		ID:        domain.ArticleID(aid),
		Title:     domain.ArticleTitle(title),
		Content:   domain.Content(content),
		Author:    domain.Author(author),
		CreatedAt: created,
		UpdatedAt: updated,
	}

	return article, nil
}

func (r *SQLArticleRepository) Create(ctx context.Context, article *domain.Article) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO articles (title, content, author, created, updated)
		VALUES (?, ?, ?, ?, ?)`,
		article.Title, article.Content, article.Author, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
