# 💈 Long parameter list

## 🥷🏻 Detection

- Your method has too many parameters.

## 💠 This Code

See `booking.go` this is an aggregate build with a lot of parameters as a constructor.
```go
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
```
We can image that Booking struct born with few parameters, and then it grows and grows with new features.
Adding code its easy, but removing or redesign is hard.


## 🧑🏻‍🔬 Refactoring

- Extract the method to a new struct and pass the struct as a parameter, this struct must group parameters using ubiquitous language.

We can see related concepts in this struct, such as `customer`, `discount`, `tax`, `date`. In some cases, we don't want class composition, but we want to group parameters, so we can use a struct to group parameters. 
And using the getters of each struct we can get the value of each parameter.

```go
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
```

But, we can go further, pushing the logic.

1. Imagine, that is necessary create logic to calculate the status of the Booking base on a Date, with the current structure are coupling to Value Object.

```go
type LocalDateTime struct {
    value time.Time
}

func (b *Booking) StatusFor(date shared.LocalDateTime) BookingStatus {
    if date.IsBefore(b.startDate) {
        return BookingStatusNotStart
    }
    if date.IsBetween(b.startDate, b.endDate) {
        return BookingStatusActive
    }
    return BookingStatusFinished
}
```

Making our data ranges aggregate meaningless, so we can push the logic to the aggregate. (See `booking_refactor.go`)
And We are not complying with tell, don´t ask principle because we are getting data to validate logic, instated of asking the object to do the logic.

```go
func (b *BookingRefactor) StatusFor(date shared.LocalDateTime) BookingStatus {
    if b.bookingPeriod.HasStarted(date) {
        return BookingStatusNotStart
    }
    if b.bookingPeriod.IsBetween(date) {
        return BookingStatusActive
    }
    return BookingStatusFinished
}
```

Appreciate the semantics!
