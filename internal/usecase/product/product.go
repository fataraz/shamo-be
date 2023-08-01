package product

import (
	"fmt"
	"log"
	productsDomain "shamo-be/internal/domain/products"
	"shamo-be/internal/shared/constant"
)

type service struct {
	productRepo productsDomain.Repository
}

func New(productRepo productsDomain.Repository) Service {
	if productRepo == nil {
		log.Fatalf("please provide product db repository")
	}
	return &service{productRepo: productRepo}
}

func (s *service) FindProducts() (resp []*ResponseProduct, err error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		err = fmt.Errorf("data not found")
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
