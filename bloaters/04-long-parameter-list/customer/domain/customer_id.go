package domain

import "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/shared/domain"

type CustomerID struct {
	value domain.UUIDValueObject
}

func (c *CustomerID) Value() string {
	return c.value.Value()
}

func NewCustomerID() (*CustomerID, error) {
	id, err := domain.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *id}, nil
}

func NewCustomerIDFromString(id string) (*CustomerID, error) {
	uid, err := domain.NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *uid}, nil
}
