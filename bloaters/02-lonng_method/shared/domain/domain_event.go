package domain

import (
	"github.com/google/uuid"
	"time"
)

type DomainEvent interface {
	GetDomainEventName() string
	GetEventId() string
	GerOccurredOn() time.Time
}

type DomainEventCommonFields struct {
	name       string
	eventId    uuid.UUID
	occurredOn time.Time
}

func (c *DomainEventCommonFields) GetDomainEventName() string {
	return c.name
}

func (c *DomainEventCommonFields) GetEventId() string {
	return c.eventId.String()
}

func (c *DomainEventCommonFields) GerOccurredOn() time.Time {
	return c.occurredOn
}

func NewDomainEventCommonFields(name string) *DomainEventCommonFields {
	return &DomainEventCommonFields{
		name:       name,
		eventId:    uuid.New(),
		occurredOn: time.Now(),
	}
}
