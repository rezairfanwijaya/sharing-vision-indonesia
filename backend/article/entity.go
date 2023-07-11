package article

import "time"

type Article struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	CreatedDate time.Time `json:"created_date"`
	UpdateDate  time.Time `json:"updated_date"`
	Status      string    `json:"status"`
}
