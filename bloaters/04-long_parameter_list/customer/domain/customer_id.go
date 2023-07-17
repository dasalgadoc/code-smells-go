package domain

import (
	shared "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/shared/domain"
)

type CustomerID struct {
	value shared.UUIDValueObject
}

func (c *CustomerID) Value() string {
	return c.value.Value()
}

func NewCustomerID() (*CustomerID, error) {
	id, err := shared.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *id}, nil
}

func NewCustomerIDFromString(id string) (*CustomerID, error) {
	uid, err := shared.NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &CustomerID{value: *uid}, nil
}
