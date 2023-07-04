package domain

type Booking struct {
	bookingID     BookingID
	startDate     LocalDateTime
	endDate       LocalDateTime
	customerId    CustomerID
	customerName  CustomerName
	customerEmail EmailAddress
	bookingType   BookingType
	discountType  DiscountType
	discountValue DiscountValue
	taxType       TaxType
	taxValue      TaxValue
}

func NewBooking(
	bookingID BookingID,
	startDate LocalDateTime,
	endDate LocalDateTime,
	customerId CustomerID,
	customerName CustomerName,
	customerEmail EmailAddress,
	bookingType BookingType,
	discountType DiscountType,
	discountValue DiscountValue,
	taxType TaxType,
	taxValue TaxValue,
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
	bookingPeriod DateRange,
	customer Customer,
	discount Discount,
	tax Tax) Booking {
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
