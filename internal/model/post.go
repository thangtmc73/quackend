package model

import "time"

type Post struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	ShortDesc   string
	Content     string
	CreatedDate time.Time
}
