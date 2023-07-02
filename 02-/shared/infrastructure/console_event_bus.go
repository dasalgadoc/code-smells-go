package infrastructure

import (
	domain2 "dasalgadoc.com/code_smell_go/02-/shared/domain"
	"fmt"
)

type ConsoleEventBus struct{}

func NewConsoleEventBus() domain2.EventBus {
	return &ConsoleEventBus{}
}

func (c *ConsoleEventBus) PublishEvent(event domain2.DomainEvent) error {
	fmt.Println("Event published: ", event.GetDomainEventName())
	return nil
}

func (c *ConsoleEventBus) PublishBulk(events []domain2.DomainEvent) error {
	for _, event := range events {
		fmt.Println("Event published: ", event.GetDomainEventName())
	}
	return nil
}
