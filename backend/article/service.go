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
	Delete(ID int) (int, error)
	GetAll(params ParamsGetAllArticles) (PaginationArticle, int, error)
}

type service struct {
	repoArticle IRepository
}

func NewService(repoArticle IRepository) *service {
	return &service{repoArticle}
}

func (s *service) Save(input InputNewArticle) (int, error) {
	// cari apakah ada title yang sama
	articleByTitle, err := s.repoArticle.FindByTitle(input.Title)
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
	isIDValid, err := helper.IsIDValid(ID)
	if !isIDValid {
		return Article{}, http.StatusBadRequest, err
	}

	articleByID, err := s.repoArticle.FindByID(ID)
	if err != nil {
		return articleByID, http.StatusInternalServerError, err
	}

	if articleByID.ID == 0 {
		return articleByID, http.StatusNotFound, fmt.Errorf("id %v tidak ditemukan", ID)
	}

	return articleByID, http.StatusOK, nil
}

func (s *service) Update(ID int, input InputNewArticle) (int, error) {
	isIDValid, err := helper.IsIDValid(ID)
	if !isIDValid {
		return http.StatusBadRequest, err
	}

	articleByID, err := s.repoArticle.FindByID(ID)
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

func (s *service) Delete(ID int) (int, error) {
	isIDValid, err := helper.IsIDValid(ID)
	if !isIDValid {
		return http.StatusBadRequest, err
	}

	articleByID, err := s.repoArticle.FindByID(ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if articleByID.ID == 0 {
		return http.StatusNotFound, fmt.Errorf("id %v tidak ditemukan", ID)
	}

	if err := s.repoArticle.Delete(ID); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *service) GetAll(params ParamsGetAllArticles) (PaginationArticle, int, error) {
	var paginataionArticle PaginationArticle

	if params.Limit < 1 {
		return paginataionArticle, http.StatusBadRequest, fmt.Errorf("limit harus lebih dari 0")
	}

	if params.Offset < 0 {
		return paginataionArticle, http.StatusBadRequest, fmt.Errorf("limit harus lebih dari sama dengan 0")
	}

	allArticles, totalData, err := s.repoArticle.FindAll(params)
	if err != nil {
		return paginataionArticle, http.StatusOK, err
	}

	// jika data kosong
	if len(allArticles) == 0 {
		return PaginationArticle{
			Limit:     params.Limit,
			TotalData: 0,
			Articles:  allArticles,
		}, http.StatusOK, nil
	}

	paginataionArticle.TotalData = totalData
	paginataionArticle.Articles = allArticles
	paginataionArticle.Limit = params.Limit

	return paginataionArticle, http.StatusOK, nil
}
