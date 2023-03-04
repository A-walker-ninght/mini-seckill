package services

import (
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/A-walker-ninght/mini-seckill/models"
)

type ProductService struct {
	productRepository _interface.Repository
}

func NewProductService(product _interface.Repository) _interface.Service {
	return &ProductService{
		productRepository: product,
	}
}

func (p *ProductService) GetByID(id int64) (models.Model, error) {
	return p.productRepository.SelectByKey(id)
}

func (p *ProductService) GetAll() ([]models.Model, error) {
	return p.productRepository.SelectAll()
}

func (p *ProductService) DeleteByID(id int64) bool {
	return p.productRepository.Delete(id)
}

func (p *ProductService) Insert(product models.Model) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *ProductService) Update(product models.Model) error {
	return p.productRepository.Update(product)
}

func (p *ProductService) GetAllInfo() (map[int]map[string]string, error) {
	return nil, nil
}
