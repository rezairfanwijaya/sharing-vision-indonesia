package handler

import (
	"net/http"
	"strconv"
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
	httpCode, err := h.articleService.Save(input)
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
		"sukses",
	)
	c.JSON(httpCode, response)
}

func (h *HandlerArticle) GetByID(c *gin.Context) {
	articleID := c.Param("id")

	// konversi ke int
	articleIDNumber, err := strconv.Atoi(articleID)
	if err != nil {
		response := helper.ResponseAPI(
			"error",
			http.StatusBadRequest,
			"id harus berupa angka dan lebih dari 0",
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	articleByID, httpCode, err := h.articleService.GetByID(articleIDNumber)
	if err != nil {
		response := helper.ResponseAPI(
			"error",
			httpCode,
			err.Error(),
		)
		c.JSON(httpCode, response)
		return
	}

	// format article
	articleByIDFormatted := article.ArticleFormatter(articleByID)

	response := helper.ResponseAPI(
		"sukses",
		httpCode,
		articleByIDFormatted,
	)
	c.JSON(httpCode, response)
}

func (h *HandlerArticle) Update(c *gin.Context) {
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

	articleID := c.Param("id")

	// konversi ke int
	articleIDNumber, err := strconv.Atoi(articleID)
	if err != nil {
		response := helper.ResponseAPI(
			"error",
			http.StatusBadRequest,
			"id harus berupa angka dan lebih dari 0",
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	httpCode, err := h.articleService.Update(articleIDNumber, input)
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
		"sukses",
	)
	c.JSON(httpCode, response)
}

func (h *HandlerArticle) Delete(c *gin.Context) {
	articleID := c.Param("id")

	// konversi ke int
	articleIDNumber, err := strconv.Atoi(articleID)
	if err != nil {
		response := helper.ResponseAPI(
			"error",
			http.StatusBadRequest,
			"id harus berupa angka dan lebih dari 0",
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// panggil service
	httpCode, err := h.articleService.Delete(articleIDNumber)
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
		"sukses",
	)
	c.JSON(httpCode, response)
}
