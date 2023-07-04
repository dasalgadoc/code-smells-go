package domain

type BookingID struct {
	value UUIDValueObject
}

func (b *BookingID) Value() string {
	return b.value.Value()
}

func NewBookingID() (*BookingID, error) {
	id, err := NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *id}, nil
}

func NewBookingIDFromString(id string) (*BookingID, error) {
	uid, err := NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &BookingID{value: *uid}, nil
}
