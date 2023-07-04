package domain

type Tax struct {
	taxType  TaxType
	taxValue TaxValue
}

func (t Tax) TaxType() TaxType {
	return t.taxType
}

func (t Tax) TaxValue() TaxValue {
	return t.taxValue
}

func NewTax(taxType TaxType, taxValue TaxValue) Tax {
	return Tax{
		taxType:  taxType,
		taxValue: taxValue,
	}
}
