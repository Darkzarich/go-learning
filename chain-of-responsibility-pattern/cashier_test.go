package main

import "testing"

func TestCashierExecute(t *testing.T) {
	c := &Cashier{}
	p := &Patient{name: "test"}
	c.execute(p) // should not panic
	if p.paymentDone {
		t.Error("paymentDone should remain false")
	}
}
