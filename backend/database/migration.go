package database

import (
	"svid/article"

	"gorm.io/gorm"
)

func migrationArticle(db *gorm.DB) error {
	var articles []article.Article

	// cek apakah sudah ada article atau belum di database
	if err := db.Find(&articles).Error; err != nil {
		return err
	}

	// jika data kosong, lakukan migrasi
	if len(articles) == 0 {
		seedArticle := article.SeedArticle()
		for _, seed := range seedArticle {
			err := db.Create(&seed).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
