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
