package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/u-masato/blogger/4-2/internal/domain"
)

func TestArticleRepository_Get(t *testing.T) {
	t.Run("記事が存在しないとエラー", func(t *testing.T) {
		repo := NewArticleRepository()
		_, err := repo.Get(1)
		assert.Error(t, err, "error should be raised")
		assert.Equal(t, "article not found by id: 1", err.Error(), "error message should be returned")
	})
	t.Run("記事が存在する", func(t *testing.T) {
		repo := NewArticleRepository()
		a := &domain.Article{
			ID:      1,
			Title:   "title",
			Content: "content",
			Author:  "author",
		}
		if err := repo.Create(a); err != nil {
			require.NoError(t, err, "error should not be raised")
		}
		actual, err := repo.Get(1)
		require.NoError(t, err, "error should not be raised")
		assert.Equal(t, a, actual, "article should be returned")
	})
}
