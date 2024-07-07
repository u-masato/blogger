package controller

import (
	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {}

func (ctrl *ArticleController) CreateArticle(c *gin.Context) {}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {}

func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {}

func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {}
