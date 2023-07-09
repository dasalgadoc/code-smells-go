package application

import (
	"dasalgadoc.com/code_smell_go/refactors/guard_clauses/domain"
	"errors"
)

// This is an awful code, but it's just an example
func Subscribe(
	user *domain.User,
	subscription *domain.Subscription,
	term domain.Term) error {
	if user != nil {
		if subscription != nil {
			if term.Value() == domain.ANUALLY {
				// subscription renew
			} else if term.Value() == domain.MONTHLY {
				// subscription renew
			} else if term.Value() == domain.QUATERLY {
				// subscription renew
			} else {
				return errors.New("term is invalid")
			}
		} else {
			return errors.New("subscription is nil")
		}
	} else {
		return errors.New("user is nil")
	}
	return nil
}

func SubscribeRefactorOne(
	user *domain.User,
	subscription *domain.Subscription,
	term domain.Term) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if subscription == nil {
		return errors.New("subscription is nil")
	}

	if term.Value() == domain.ANUALLY {
	}

	if term.Value() == domain.MONTHLY {
	}

	if term.Value() == domain.QUATERLY {
	}

	return errors.New("term is invalid")
}
