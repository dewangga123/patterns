package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must Exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using debit") {
		t.Error("The debit payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}
