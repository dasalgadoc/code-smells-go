package domain

type Subscription struct{}

func (s *Subscription) IsExpired() bool {
	return false
}
