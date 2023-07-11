package router

import (
	"svid/article"
	"svid/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, dbConnection *gorm.DB) {
	repoArticle := article.NewRepository(dbConnection)
	serviceArticle := article.NewService(repoArticle)
	handlerArticle := handler.NewHandlerArticle(serviceArticle)

	router.POST("/article", handlerArticle.Save)
}
