package article

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Save(article Article) error
	FindByID(ID int) (Article, error)
	FindByTitle(title string) (Article, error)
	FindAll(params ParamsGetAllArticles) ([]Article, int, error)
	Update(article Article) error
	Delete(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(article Article) error {
	if err := r.db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(ID int) (Article, error) {
	var article Article
	if err := r.db.Where("id = ?", ID).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) FindByTitle(title string) (Article, error) {
	var article Article
	if err := r.db.Where("title = ?", title).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) FindAll(params ParamsGetAllArticles) ([]Article, int, error) {
	var articles []Article
	var totalData int64 = 0

	if err := r.db.Limit(params.Limit).Offset(params.Offset).Order("id DESC").Find(&articles).Error; err != nil {
		return articles, int(totalData), err
	}

	// total data
	if err := r.db.Model(&Article{}).Count(&totalData).Error; err != nil {
		return articles, int(totalData), err
	}

	return articles, int(totalData), nil
}

func (r *repository) Update(article Article) error {
	if err := r.db.Save(&article).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ID int) error {
	var article Article
	if err := r.db.Where("id = ?", ID).Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
