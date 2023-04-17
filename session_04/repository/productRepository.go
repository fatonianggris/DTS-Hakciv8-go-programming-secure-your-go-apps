package repository

import "go-programming-secure-your-go-apps/session_04/entity"

type ProductRepository interface {
	FindById(id int) *entity.Product
	FindAll() *[]entity.Product
}
