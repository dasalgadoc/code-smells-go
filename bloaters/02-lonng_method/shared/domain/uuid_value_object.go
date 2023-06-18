package domain

import (
	"github.com/google/uuid"
)

type UUIDValueObject struct {
	value uuid.UUID
}

func NewUUIDValueObject() (*UUIDValueObject, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &UUIDValueObject{value: id}, err
}

func NewUUIDValueObjectFromString(id string) (*UUIDValueObject, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &UUIDValueObject{value: uid}, err
}

func (s *UUIDValueObject) Value() string {
	return s.value.String()
}
