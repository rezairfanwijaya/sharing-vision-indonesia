package article

import "gorm.io/gorm"

type IRepository interface {
	Save()
	GetByID()
	Update()
	Delete()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
