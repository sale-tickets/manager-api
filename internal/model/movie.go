package model

import "github.com/lib/pq"

type Movie struct {
	Base
	CategoryID  pq.StringArray `json:"category_id" gorm:"type:text[]"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Thumbnail   string         `json:"thumbnail"`
}
