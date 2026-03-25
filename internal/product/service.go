package product

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) GetProduct(id uint64) (Product, error) {
	return s.repo.GetByID(id)
}

func (s Service) ListProducts() []Product {
	return s.repo.List()
}
