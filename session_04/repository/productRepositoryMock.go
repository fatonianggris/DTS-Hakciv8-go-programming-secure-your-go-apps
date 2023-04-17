package repository

import (
	"go-programming-secure-your-go-apps/session_04/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (productRepositoryMock ProductRepositoryMock) FindById(id int) *entity.Product {
	arguments := productRepositoryMock.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)

	return &product
}

func (productRepositoryMock ProductRepositoryMock) FindAll() *[]entity.Product {
	arguments := productRepositoryMock.Mock.MethodCalled("FindAll")
	if arguments.Get(0) == nil {
		return nil
	}

	products := arguments.Get(0).([]entity.Product)

	return &products
}
