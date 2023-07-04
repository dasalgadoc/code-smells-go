package domain

import "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/shared/domain"

type BookingID struct {
	value domain.UUIDValueObject
}

func (b *BookingID) Value() string {
	return b.value.Value()
}

func NewBookingID() (*BookingID, error) {
	id, err := domain.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *id}, nil
}

func NewBookingIDFromString(id string) (*BookingID, error) {
	uid, err := domain.NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *uid}, nil
}