package products

type UseCase interface {
	SearchProducts(keyword string) ([]*Products, error)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

type useCase struct {
	repo Repository
}

// SearchProducts implements UseCase
func (uc *useCase) SearchProducts(keyword string) (results []*Products, err error) {
	return uc.repo.FindByKeyword(keyword)
}
