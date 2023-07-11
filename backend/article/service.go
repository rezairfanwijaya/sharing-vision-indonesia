package article

import (
	"fmt"
	"net/http"
	"svid/helper"
	"time"
)

type IService interface {
	Save(input InputNewArticle) (Article, int, error)
	// GetByID(ID int) (Article, error)
	// GetAll() ([]Article, error)
	// Update(ID int) error
	// Delete(ID int) error
}

type service struct {
	repoArticle IRepository
}

func NewService(repoArticle IRepository) *service {
	return &service{repoArticle}
}

func (s *service) Save(input InputNewArticle) (Article, int, error) {
	// cari apakah ada title yang sama
	articleByTitle, err := s.repoArticle.GetByTitle(input.Title)
	if err != nil {
		return articleByTitle, http.StatusInternalServerError, err
	}

	// tidak boleh ada judul yang sama
	if articleByTitle.Title == input.Title {
		return articleByTitle, http.StatusConflict, fmt.Errorf("judul %v sudah dipakai", input.Title)
	}

	// validasi status
	isStatusValid := helper.IsStatusValid(input.Status)
	if !isStatusValid {
		return articleByTitle, http.StatusBadRequest, fmt.Errorf("status yang didukung hanya publish, draft atau thrash")
	}

	// binding
	article := Article{
		Title:       input.Title,
		Content:     input.Content,
		Category:    input.Category,
		CreatedDate: time.Now(),
		Status:      input.Status,
	}

	// save
	newArticle, err := s.repoArticle.Save(article)
	if err != nil {
		return newArticle, http.StatusInternalServerError, err
	}

	return newArticle, http.StatusOK, nil
}
