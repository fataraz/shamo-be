package products

// Repository ...
type Repository interface {
	FindAll() (resp []*Product, err error)
}
