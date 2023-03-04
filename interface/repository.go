package _interface

import "github.com/A-walker-ninght/mini-seckill/models"

type Repository interface {
	Conn() error
	Insert(order models.Model) (int64, error)
	Delete(int64) bool
	Update(order models.Model) error
	SelectByKey(int64) (models.Model, error)
	SelectAll() ([]models.Model, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}
