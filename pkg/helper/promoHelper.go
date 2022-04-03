package helper

func PromoHelper(PromoCode string, Price int64, Quantity int64) (int64, error) {
	switch PromoCode {
	case "BELI2GRATIS1":
		if Quantity > 2 {
			return Price * (Quantity - 1), nil
		}
		return Price - (Price * 10 / 100), nil
	case "BELI3GRATIS1":
		if Quantity > 3 {
			return Price * (Quantity - 1), nil
		}
		return Price - (Price * 20 / 100), nil
	case "DISKON10%":
		TotalPrice := Price * Quantity
		return TotalPrice - (TotalPrice * 10 / 100), nil
	default:
		TotalPrice := Price * Quantity
		return TotalPrice, nil
	}
}
