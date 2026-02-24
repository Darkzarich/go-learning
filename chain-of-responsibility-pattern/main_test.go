package main

import "testing"

func TestFullChain(t *testing.T) {
	cashier := &Cashier{}
	medical := &Medical{}
	medical.setNext(cashier)
	doctor := &Doctor{}
	doctor.setNext(medical)
	reception := &Reception{}
	reception.setNext(doctor)
	p := &Patient{name: "test"}
	reception.execute(p)
	if !p.registrationDone {
		t.Error("registration not done")
	}
	if !p.doctorCheckUpDone {
		t.Error("doctor check not done")
	}
	if !p.medicineDone {
		t.Error("medicine not done")
	}
	if p.paymentDone {
		t.Error("payment should not be done")
	}
}
