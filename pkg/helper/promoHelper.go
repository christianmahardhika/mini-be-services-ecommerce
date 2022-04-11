package helper

func PromoHelper(Discount int64, Price int64, Quantity int64) (int64, error) {
	return Price - (Price * (Discount / 100)), nil
}
