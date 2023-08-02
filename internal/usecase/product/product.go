package product

import (
	"log"
	ctxSess "shamo-be/internal/shared/utils/context"

	productsDomain "shamo-be/internal/domain/products"
	"shamo-be/internal/shared/constant"
)

// service ...
type service struct {
	productRepo productsDomain.Repository
}

// New ...
func New(productRepo productsDomain.Repository) Service {
	if productRepo == nil {
		log.Fatalf("please provide product db repository")
	}
	return &service{productRepo: productRepo}
}

// FindProducts ...
func (s *service) FindProducts(ctxSess *ctxSess.Context) (resp []*ResponseProduct, err error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constant.ErrorDataNotFound
		return
	}
	for _, v := range products {
		product := &ResponseProduct{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			CategoryID:  v.CategoryID,
			tags:        v.Tags,
			CreatedAt:   v.CreatedAt.Format(constant.FormatDateString),
			UpdatedAt:   v.UpdatedAt.Format(constant.FormatDateString),
		}
		resp = append(resp, product)
	}
	return
}
