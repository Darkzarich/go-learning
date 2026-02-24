package main

import "testing"

func TestDoctorExecute_WhenNotChecked(t *testing.T) {
	mock := &mockDepartment{}
	d := &Doctor{}
	d.setNext(mock)
	p := &Patient{name: "test"}
	d.execute(p)
	if !p.doctorCheckUpDone {
		t.Error("doctorCheckUpDone not set")
	}
	if !mock.called {
		t.Error("next not called")
	}
	if mock.p != p {
		t.Error("wrong patient passed")
	}
}

func TestDoctorExecute_WhenAlreadyChecked(t *testing.T) {
	mock := &mockDepartment{}
	d := &Doctor{}
	d.setNext(mock)
	p := &Patient{name: "test", doctorCheckUpDone: true}
	d.execute(p)
	if !mock.called {
		t.Error("next not called even when checked")
	}
}
