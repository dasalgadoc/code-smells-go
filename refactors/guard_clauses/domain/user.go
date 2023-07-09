package domain

type User struct{}

func (u *User) IsSubscribed() bool {
	return false
}

func (u *User) Subscribe(term *Term) error {
	return nil
}
