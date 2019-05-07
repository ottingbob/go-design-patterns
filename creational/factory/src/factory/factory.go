package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash	  = 1
	DebitCard = 2
	// New implementation
	CreditCard = 3
)

type CashPM struct{}
type DebitCardPM struct{}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

// New implementation to add
type CreditCardPM struct {}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n", amount)
}