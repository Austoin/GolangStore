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

func (s Service) ListCheckedItems(userID uint64) []Item {
	items := s.repo.ListByUserID(userID)
	checked := make([]Item, 0)
	for _, item := range items {
		if item.Checked {
			checked = append(checked, item)
		}
	}

	return checked
}
