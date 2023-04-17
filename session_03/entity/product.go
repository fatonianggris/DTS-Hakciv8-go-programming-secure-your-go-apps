package entity

import "time"

type Product struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null;type:varchar(255)"`
	Description string `json:"description" gorm:"not null;type:varchar(255)"`
	UserId      uint
	User        *User
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
