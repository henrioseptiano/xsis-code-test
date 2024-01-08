package model

import "time"

type Movie struct {
	ID          int64      `json:"id" gorm:"id,primaryKey,autoIncrement"`
	Title       string     `json:"title" gorm:"title,not null"`
	Description string     `json:"description" gorm:"description,not null"`
	Rating      float32    `json:"rating" gorm:"rating, not null"`
	Image       string     `json:"image" gorm:"image, not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at,not null"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"updated_at,not null"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"deleted_at"`
}
