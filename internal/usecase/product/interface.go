package product

// Service ...
type Service interface {
	FindProducts() (resp []*ResponseProduct, err error)
}
