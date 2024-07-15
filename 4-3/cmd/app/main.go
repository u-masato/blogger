package main

import (
	"github.com/gin-gonic/gin"
	"github.com/u-masato/blogger/4-3/internal/controller"
	"github.com/u-masato/blogger/4-3/internal/infra"
	"github.com/u-masato/blogger/4-3/internal/repository"
	"github.com/u-masato/blogger/4-3/internal/usecase"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// データベース初期化
	db := infra.InitDB()
	tx := infra.NewTxAdmin(db)

	articleRepo := repository.NewSQLArticleRepository(db)
	articleUC := usecase.NewArticleUsecase(articleRepo, tx)
	articleController := controller.NewArticleController(articleUC)

	// 記事関連のエンドポイント
	articleRoutes := router.Group("/articles")
	{
		articleRoutes.GET("/", articleController.GetAllArticles)      // 全記事取得
		articleRoutes.POST("/", articleController.CreateArticle)      // 記事作成
		articleRoutes.GET("/:id", articleController.GetArticle)       // 記事取得
		articleRoutes.PUT("/:id", articleController.UpdateArticle)    // 記事更新
		articleRoutes.DELETE("/:id", articleController.DeleteArticle) // 記事削除
	}
	return router
}

// go run main.go
func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
