package domain

import "time"

type UserRefactor struct {
	id        UserId
	email     UserEmail
	birthDate UserBirthDate
}

// Constructor with primitive types, Value objects and modeling are not exposed
func NewUserVo(id, email string, birthDate time.Time) (*UserRefactor, error) {
	userId, err := NewUserId(id)
	if err != nil {
		return nil, err
	}
	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}
	userBirthDate, err := NewUserBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	return &UserRefactor{
		id:        *userId,
		email:     *userEmail,
		birthDate: *userBirthDate,
	}, nil
}

func (u *UserRefactor) UpdateEmail(newEmail string) error {
	userEmail, err := NewUserEmail(newEmail)
	if err != nil {
		return err
	}
	u.email = *userEmail
	return nil
}

/*--Getters--*/
func (u *UserRefactor) Id() string {
	return u.id.Value()
}

func (u *UserRefactor) Email() string {
	return u.email.Value()
}

func (u *UserRefactor) BirthDate() time.Time {
	return u.birthDate.Value()
}
