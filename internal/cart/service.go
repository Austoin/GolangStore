package cart

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) ListItems(userID uint64) []Item {
	return s.repo.ListByUserID(userID)
}
