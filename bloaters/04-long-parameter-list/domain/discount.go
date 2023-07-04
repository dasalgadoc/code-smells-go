package domain

type Discount struct {
	discountType  DiscountType
	discountValue DiscountValue
}

func (d Discount) DiscountType() DiscountType {
	return d.discountType
}

func (d Discount) DiscountValue() DiscountValue {
	return d.discountValue
}

func NewDiscount(discountType DiscountType, discountValue DiscountValue) Discount {
	return Discount{
		discountType:  discountType,
		discountValue: discountValue,
	}
}
