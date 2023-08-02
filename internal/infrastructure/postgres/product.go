package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
	productsDomain "shamo-be/internal/domain/products"
	"shamo-be/internal/shared/constant"
	"shamo-be/internal/shared/database"
)

// product ...
type product struct {
	db *database.Database
}

// NewProductsRepository ...
func NewProductsRepository(db *database.Database) productsDomain.Repository {
	if db == nil {
		panic("please provide database client")
	}
	return &product{db: db}
}

// Save ...
func (p *product) Save(ctx context.Context, entity *productsDomain.Product) (err error) {
	return
}

// FindAll ...
func (p *product) FindAll() (resp []*productsDomain.Product, err error) {
	resp = []*productsDomain.Product{}
	err = p.db.Find(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constant.ErrorProductNotFound
			return
		}
		err = constant.ErrorDatabase
		return
	}
	return
}
