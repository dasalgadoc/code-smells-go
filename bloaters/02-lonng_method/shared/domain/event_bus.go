package domain

type EventBus interface {
	PublishEvent(event DomainEvent) error
	PublishBulk(events []DomainEvent) error
}
