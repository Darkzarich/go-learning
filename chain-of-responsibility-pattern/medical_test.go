package main

import "testing"

func TestMedicalExecute_WhenNotGiven(t *testing.T) {
	mock := &mockDepartment{}
	m := &Medical{}
	m.setNext(mock)
	p := &Patient{name: "test"}
	m.execute(p)
	if !p.medicineDone {
		t.Error("medicineDone not set")
	}
	if !mock.called {
		t.Error("next not called")
	}
	if mock.p != p {
		t.Error("wrong patient passed")
	}
}

func TestMedicalExecute_WhenAlreadyGiven(t *testing.T) {
	mock := &mockDepartment{}
	m := &Medical{}
	m.setNext(mock)
	p := &Patient{name: "test", medicineDone: true}
	m.execute(p)
	if !mock.called {
		t.Error("next not called even when given")
	}
}
