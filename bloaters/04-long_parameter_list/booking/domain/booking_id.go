package domain

import (
	shared "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/shared/domain"
)

type BookingID struct {
	value shared.UUIDValueObject
}

func (b *BookingID) Value() string {
	return b.value.Value()
}

func NewBookingID() (*BookingID, error) {
	id, err := shared.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *id}, nil
}

func NewBookingIDFromString(id string) (*BookingID, error) {
	uid, err := shared.NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *uid}, nil
}
