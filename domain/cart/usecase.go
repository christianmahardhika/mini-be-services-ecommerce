package cart

type UseCase interface {
	GetCartInfo() ([]Cart, error)
	AddItemToCart(item Cart) (string, error)
	DeleteItemFromCart(id int64) (string, error)
	ResetCart() (string, error)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

type useCase struct {
	repo Repository
}

// AddItemToCart implements UseCase
func (uc *useCase) AddItemToCart(item Cart) (string, error) {
	err := uc.repo.UpdateInsert(&item)
	if err != nil {
		return "", err
	}

	return "item added to cart", nil
}

// DeleteItemFromCart implements UseCase
func (uc *useCase) DeleteItemFromCart(id int64) (string, error) {
	err := uc.repo.Delete(id)
	if err != nil {
		return "Cannot Delete Item / Item has empty", err
	}

	return "item deleted", nil
}

// GetCartInfo implements UseCase
func (uc *useCase) GetCartInfo() ([]Cart, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil

}

// ResetCart implements UseCase
func (uc *useCase) ResetCart() (string, error) {
	err := uc.repo.DeleteAll()
	if err != nil {
		return "", err
	}
	return "cart has been reset", nil
}
