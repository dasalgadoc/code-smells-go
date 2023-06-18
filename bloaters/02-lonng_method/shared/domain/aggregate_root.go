package domain

type AggregateRoot struct {
	eventList []DomainEvent
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		eventList: make([]DomainEvent, 0),
	}
}

func (a *AggregateRoot) PullEvents() []DomainEvent {
	events := a.eventList
	a.eventList = make([]DomainEvent, 0)

	return events
}

func (a *AggregateRoot) RecordEvent(event DomainEvent) {
	a.eventList = append(a.eventList, event)
}
