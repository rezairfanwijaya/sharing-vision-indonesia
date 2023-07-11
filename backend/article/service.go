package article

import (
	"fmt"
	"net/http"
	"svid/helper"
	"time"
)

type IService interface {
	Save(input InputNewArticle) (int, error)
	GetByID(ID int) (Article, int, error)
	Update(ID int, input InputNewArticle) (int, error)
	// GetAll() ([]Article, error)
	// Delete(ID int) error
}

type service struct {
	repoArticle IRepository
}

func NewService(repoArticle IRepository) *service {
	return &service{repoArticle}
}

func (s *service) Save(input InputNewArticle) (int, error) {
	// cari apakah ada title yang sama
	articleByTitle, err := s.repoArticle.GetByTitle(input.Title)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// tidak boleh ada judul yang sama
	if articleByTitle.Title == input.Title {
		return http.StatusConflict, fmt.Errorf("judul %v sudah dipakai", input.Title)
	}

	// validasi status
	isStatusValid := helper.IsStatusValid(input.Status)
	if !isStatusValid {
		return http.StatusBadRequest, fmt.Errorf("status yang didukung hanya publish, draft atau thrash")
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
	err = s.repoArticle.Save(article)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *service) GetByID(ID int) (Article, int, error) {
	articleByID, err := s.repoArticle.GetByID(ID)
	if err != nil {
		return articleByID, http.StatusInternalServerError, err
	}

	if articleByID.ID == 0 {
		return articleByID, http.StatusNotFound, fmt.Errorf("id %v tidak ditemukan", ID)
	}

	return articleByID, http.StatusOK, nil
}

func (s *service) Update(ID int, input InputNewArticle) (int, error) {
	articleByID, err := s.repoArticle.GetByID(ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if articleByID.ID == 0 {
		return http.StatusNotFound, fmt.Errorf("id %v tidak ditemukan", ID)
	}

	// validasi status
	isStatusValid := helper.IsStatusValid(input.Status)
	if !isStatusValid {
		return http.StatusBadRequest, fmt.Errorf("status yang didukung hanya publish, draft atau thrash")
	}

	// bind
	articleByID.Category = input.Category
	articleByID.Content = input.Content
	articleByID.Status = input.Status
	articleByID.Title = input.Title
	articleByID.UpdateDate = time.Now()

	// panggil repo
	if err := s.repoArticle.Update(articleByID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
