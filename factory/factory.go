package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

const (
	Cash      = 1
	DebitCard = 2
)

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return "Paid with cash."
}

func (d *DebitCardPM) Pay(amount float64) string {
	return "Paid with debit card."
}
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognizes\n", m))
	}
}
