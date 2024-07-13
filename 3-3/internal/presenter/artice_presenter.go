package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/u-masato/blogger/3-3/internal/domain"
)

type ArticlePresenter struct {
}

func NewArticleCreatePresenter() *ArticlePresenter {
	return &ArticlePresenter{}
}

func (p *ArticlePresenter) Progress(percentage int) {
	// ユーザーには表示されない
	fmt.Printf("進捗... %d%%\n", percentage)
}

func (p *ArticlePresenter) Complete() {
	// ユーザーには表示されない
	fmt.Println("完了")
}

func (p *ArticlePresenter) Present(article *domain.Article) {
	a := map[string]interface{}{
		"ID":      article.ID,
		"Title":   article.Title,
		"Content": article.Content,
		"Author":  article.Author,
		// 時刻を文字列に変換
		"Created": article.CreatedAt.Format("2006-01-02 15:04"),
		"Updated": article.UpdatedAt.Format("2006-01-02 15:04"),
	}
	// マップをJSON形式にエンコード
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// JSONを標準出力に表示
	// ユーザーには表示されない
	fmt.Println(string(jsonData))
}
