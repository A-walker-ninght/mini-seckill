package _interface

import "github.com/A-walker-ninght/mini-seckill/models"

type Service interface {
	GetByID(int64) (models.Model, error)
	DeleteByID(int64) bool
	Update(models.Model) error
	Insert(model models.Model) (int64, error)
	GetAll() ([]models.Model, error)
	GetAllInfo() (map[int]map[string]string, error)
}
