package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// 記事関連のエンドポイント
	articleRoutes := router.Group("/articles")
	{
		articleRoutes.GET("/", getAllArticles)      // 全記事取得
		articleRoutes.POST("/", createArticle)      // 記事作成
		articleRoutes.GET("/:id", getArticle)       // 記事取得
		articleRoutes.PUT("/:id", updateArticle)    // 記事更新
		articleRoutes.DELETE("/:id", deleteArticle) // 記事削除
	}
	return router
}

// 記事関連のハンドラー（未実装）
func getAllArticles(c *gin.Context) {}
func createArticle(c *gin.Context)  {}
func getArticle(c *gin.Context)     {}
func updateArticle(c *gin.Context)  {}
func deleteArticle(c *gin.Context)  {}

// go run main.go
func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
