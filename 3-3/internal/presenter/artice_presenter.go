package presenter

/*
	プレゼンターは表示を担当します
	ユースケース層からの出力をユーザーに表示する

	クリーンアーキテクチャの制御の流れの図の通り、
	Output Portはusecaseパッケージに定義されているインターフェースで、
	プレゼンターはOutput Portの実装です
	プレゼンターはユースケースに注入していますが、プレゼンターがユースケースをし呼び出す方法もあります

	ただし、MVCとは相性が悪く
	https://nrslib.com/clean-architecture/#outline__5_1_3 で説明されています
	デスクトップアプリケーションなどは、ユースケースの処理中に表示ロジックを挟むメリットはありますが、
	Webアプリケーションフレームワークでは、コントローラーがレスポンスを返すため、ユースケースの処理が終わってから表示ロジックを挟む方が一般的です（WebSocket などの非同期でサーバープッシュするような仕組みで動作するようなシステムでなければ）

	そのため、JSON　APIのような簡単な変換処理の場合は、プレゼンターをユースケースに注入する必要はなく、
	コントローラーで直接変換処理を行うか、プレゼンターを呼び出すこともできます
	この場合、コントローラーがプレゼンターに依存することになりますが、依存の方向に違反しているわけではありません
*/
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
