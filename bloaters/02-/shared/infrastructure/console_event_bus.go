package infrastructure

import (
	"dasalgadoc.com/code_smell_go/bloaters/02-/shared/domain"
	"fmt"
)

type ConsoleEventBus struct{}

func NewConsoleEventBus() domain.EventBus {
	return &ConsoleEventBus{}
}

func (c *ConsoleEventBus) PublishEvent(event domain.DomainEvent) error {
	fmt.Println("Event published: ", event.GetDomainEventName())
	return nil
}

func (c *ConsoleEventBus) PublishBulk(events []domain.DomainEvent) error {
	for _, event := range events {
		fmt.Println("Event published: ", event.GetDomainEventName())
	}
	return nil
}
