package entity

import (
	"go-programming-secure-your-go-apps/session_03/helper"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Email     string    `json:"email" gorm:"not null;unique;type:varchar(50)"`
	Password  string    `json:"password" gorm:"not null;type:varchar(255)"`
	Products  []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = helper.HashPassword(user.Password)
	err = nil
	return
}
