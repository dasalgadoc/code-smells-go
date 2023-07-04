package domain

import (
	"dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/customer/domain"
	domain3 "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/discounts/domain"
	domain2 "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/shared/domain"
	domain4 "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/taxes/domain"
)

type Booking struct {
	bookingID     BookingID
	startDate     domain2.LocalDateTime
	endDate       domain2.LocalDateTime
	customerId    domain.CustomerID
	customerName  domain.CustomerName
	customerEmail domain.EmailAddress
	bookingType   BookingType
	discountType  domain3.DiscountType
	discountValue domain3.DiscountValue
	taxType       domain4.TaxType
	taxValue      domain4.TaxValue
}

func NewBooking(
	bookingID BookingID,
	startDate domain2.LocalDateTime,
	endDate domain2.LocalDateTime,
	customerId domain.CustomerID,
	customerName domain.CustomerName,
	customerEmail domain.EmailAddress,
	bookingType BookingType,
	discountType domain3.DiscountType,
	discountValue domain3.DiscountValue,
	taxType domain4.TaxType,
	taxValue domain4.TaxValue,
) Booking {
	return Booking{
		bookingID:     bookingID,
		startDate:     startDate,
		endDate:       endDate,
		customerId:    customerId,
		customerName:  customerName,
		customerEmail: customerEmail,
		bookingType:   bookingType,
		discountType:  discountType,
		discountValue: discountValue,
		taxType:       taxType,
		taxValue:      taxValue,
	}
}

func NewBookingRefactor(
	bookingID BookingID,
	bookingType BookingType,
	bookingPeriod domain2.DateRange,
	customer domain.Customer,
	discount domain3.Discount,
	tax domain4.Tax) Booking {
	return Booking{
		bookingID:     bookingID,
		startDate:     bookingPeriod.StartDate(),
		endDate:       bookingPeriod.EndDate(),
		bookingType:   bookingType,
		customerId:    customer.CustomerId(),
		customerName:  customer.CustomerName(),
		customerEmail: customer.CustomerEmail(),
		discountType:  discount.DiscountType(),
		discountValue: discount.DiscountValue(),
		taxType:       tax.TaxType(),
		taxValue:      tax.TaxValue(),
	}
}
