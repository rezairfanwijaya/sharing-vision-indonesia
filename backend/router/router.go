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
	router.GET("/article/:id", handlerArticle.GetByID)
	router.PUT("/article/:id", handlerArticle.Update)
	router.DELETE("/article/:id", handlerArticle.Delete)
	router.GET("/articles/:limit/:offset", handlerArticle.GetAll)
}
