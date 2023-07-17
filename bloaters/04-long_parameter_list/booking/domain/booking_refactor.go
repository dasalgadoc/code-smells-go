package domain

import (
	customers "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/customer/domain"
	discounts "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/discounts/domain"
	shared "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/shared/domain"
	taxes "dasalgadoc.com/code_smell_go/bloaters/04-long_parameter_list/taxes/domain"
)

type BookingRefactor struct {
	bookingID     BookingID
	bookingPeriod shared.DateRange
	bookingType   BookingType
	customer      customers.Customer
	discount      discounts.Discount
	tax           taxes.Tax
}

func NewBookingRefactorStruct(
	bookingID BookingID,
	bookingPeriod shared.DateRange,
	bookingType BookingType,
	customer customers.Customer,
	discount discounts.Discount,
	tax taxes.Tax) BookingRefactor {
	return BookingRefactor{
		bookingID:     bookingID,
		bookingPeriod: bookingPeriod,
		bookingType:   bookingType,
		customer:      customer,
		discount:      discount,
		tax:           tax,
	}
}

func (b *BookingRefactor) StatusFor(date shared.LocalDateTime) BookingStatus {
	if b.bookingPeriod.HasStarted(date) {
		return BookingStatusNotStart
	}
	if b.bookingPeriod.IsBetween(date) {
		return BookingStatusActive
	}
	return BookingStatusFinished
}
