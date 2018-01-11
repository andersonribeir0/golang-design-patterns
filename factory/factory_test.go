package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.1)
	if !strings.Contains(msg, "cash") {
		t.Error("The cash payment method message wasn't correct.")
	}
	t.Log("Log: ", msg)
}
