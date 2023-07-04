package domain

type Customer struct {
	customerId    CustomerID
	customerName  CustomerName
	customerEmail EmailAddress
}

func (c Customer) CustomerId() CustomerID {
	return c.customerId
}

func (c Customer) CustomerName() CustomerName {
	return c.customerName
}

func (c Customer) CustomerEmail() EmailAddress {
	return c.customerEmail
}

func NewCustomer(
	customerId CustomerID,
	customerName CustomerName,
	customerEmail EmailAddress,
) Customer {
	return Customer{
		customerId:    customerId,
		customerName:  customerName,
		customerEmail: customerEmail,
	}
}
