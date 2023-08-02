package product

import ctxSess "shamo-be/internal/shared/utils/context"

// Service ...
type Service interface {
	FindProducts(ctxSess *ctxSess.Context) (resp []*ResponseProduct, err error)
}
