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
