package handler

import (
	"net/http"
	"svid/article"
	"svid/helper"

	"github.com/gin-gonic/gin"
)

type HandlerArticle struct {
	articleService article.IService
}

func NewHandlerArticle(articleService article.IService) *HandlerArticle {
	return &HandlerArticle{articleService}
}

func (h *HandlerArticle) Save(c *gin.Context) {
	var input article.InputNewArticle

	// bind
	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.ResponseAPI(
			"error binding",
			http.StatusBadRequest,
			errBinding,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	newArticle, httpCode, err := h.articleService.Save(input)
	if err != nil {
		response := helper.ResponseAPI(
			"error",
			httpCode,
			err.Error(),
		)
		c.JSON(httpCode, response)
		return
	}

	response := helper.ResponseAPI(
		"sukses",
		httpCode,
		newArticle,
	)
	c.JSON(httpCode, response)
}
