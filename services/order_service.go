package services

import (
	_interface "github.com/A-walker-ninght/mini-seckill/interface"
	"github.com/A-walker-ninght/mini-seckill/models"
)

type OrderService struct {
	OrderRepository _interface.Repository
}

func NewOrderService(repository _interface.Repository) _interface.Service {
	return &OrderService{
		OrderRepository: repository,
	}
}

func (o *OrderService) GetByID(i int64) (models.Model, error) {
	return o.OrderRepository.SelectByKey(i)
}

func (o *OrderService) DeleteByID(i int64) bool {
	return o.OrderRepository.Delete(i)
}

func (o *OrderService) Update(order models.Model) error {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) Insert(order models.Model) (int64, error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) GetAll() ([]models.Model, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllInfo() (map[int]map[string]string, error) {
	return o.OrderRepository.SelectAllWithInfo()
}
