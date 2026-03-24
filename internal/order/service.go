package order

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) GetOrder(orderNo string) (Order, error) {
	return s.repo.GetByOrderNo(orderNo)
}
