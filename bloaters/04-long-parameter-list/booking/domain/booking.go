package domain

import (
	customers "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/customer/domain"
	discounts "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/discounts/domain"
	shared "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/shared/domain"
	taxes "dasalgadoc.com/code_smell_go/bloaters/04-long-parameter-list/taxes/domain"
)

type Booking struct {
	bookingID     BookingID
	startDate     shared.LocalDateTime
	endDate       shared.LocalDateTime
	customerId    customers.CustomerID
	customerName  customers.CustomerName
	customerEmail customers.EmailAddress
	bookingType   BookingType
	discountType  discounts.DiscountType
	discountValue discounts.DiscountValue
	taxType       taxes.TaxType
	taxValue      taxes.TaxValue
}

func NewBooking(
	bookingID BookingID,
	startDate shared.LocalDateTime,
	endDate shared.LocalDateTime,
	customerId customers.CustomerID,
	customerName customers.CustomerName,
	customerEmail customers.EmailAddress,
	bookingType BookingType,
	discountType discounts.DiscountType,
	discountValue discounts.DiscountValue,
	taxType taxes.TaxType,
	taxValue taxes.TaxValue,
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
	bookingPeriod shared.DateRange,
	customer customers.Customer,
	discount discounts.Discount,
	tax taxes.Tax) Booking {
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
