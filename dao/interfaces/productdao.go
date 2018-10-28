package interfaces

import "cqrs-test/model"

// ProductDao interface
type ProductDao interface {
	Create(u *model.Product) error
	GetByID(i string) (model.Product, error)
	GetAll() ([]model.Product, error)
}
