package repository

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IRepository[TData any] interface {
	Insert(data *TData) error
	Update(data *TData) error
	Delete(id string) error
	Get(id string) (*TData, error)
	GetAll() ([]TData, error)
}

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) IRepository[T] {
	return Repository[T]{db: db}
}

func (r Repository[TData]) Insert(data *TData) error {
	result := r.db.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository[TData]) Update(data *TData) error {
	return r.db.Save(data).Error
}

func (r Repository[TData]) Delete(id string) error {
	var result TData
	return r.db.Delete(&result, id).Error
}

func (r Repository[TData]) Get(id string) (*TData, error) {
	return r.Get(id)
}

func (r Repository[TData]) GetAll() ([]TData, error) {
	var result []TData
	return result, r.db.Find(&result).Error
}
