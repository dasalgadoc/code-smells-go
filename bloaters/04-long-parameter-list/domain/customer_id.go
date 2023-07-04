package domain

type CustomerID struct {
	value UUIDValueObject
}

func (c *CustomerID) Value() string {
	return c.value.Value()
}

func NewCustomerID() (*CustomerID, error) {
	id, err := NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *id}, nil
}

func NewCustomerIDFromString(id string) (*CustomerID, error) {
	uid, err := NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *uid}, nil
}
