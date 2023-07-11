package article

import "gorm.io/gorm"

type IRepository interface {
	Save(article Article) error
	GetByID(ID int) (Article, error)
	GetByTitle(title string) (Article, error)
	GetAll() ([]Article, error)
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

func (r *repository) GetByID(ID int) (Article, error) {
	var article Article
	if err := r.db.Where("id = ?", ID).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) GetByTitle(title string) (Article, error) {
	var article Article
	if err := r.db.Where("title = ?", title).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) GetAll() ([]Article, error) {
	var articles []Article
	if err := r.db.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
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
