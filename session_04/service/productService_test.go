package service

import (
	"go-programming-secure-your-go-apps/session_04/entity"
	"go-programming-secure-your-go-apps/session_04/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductFound(t *testing.T) {
	product := entity.Product{
		Id:   2,
		Name: "Camera",
	}

	productRepository.Mock.On("FindById", 2).Return(product)
	result := productService.GetOneProduct(2)

	assert.Equal(t, &product, result)
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", 1).Return(nil)
	result := productService.GetOneProduct(1)

	assert.Nil(t, result)
}

func TestProductServiceGetAllProductsFound(t *testing.T) {
	products := []entity.Product{
		{Id: 1, Name: "Phone"},
		{Id: 2, Name: "Camera"},
	}

	productRepository.Mock.On("FindAll").Return(products).Once()
	result := productService.GetAllProducts()

	assert.Equal(t, &products, result)
}

func TestProductServiceGetAllProductsNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil).Once()
	result := productService.GetAllProducts()

	assert.Nil(t, result)
}
