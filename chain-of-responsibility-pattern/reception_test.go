package main

import "testing"

type mockDepartment struct {
	called bool
	p      *Patient
}

func (m *mockDepartment) execute(p *Patient) {
	m.called = true
	m.p = p
}

func (m *mockDepartment) setNext(Department) {}

func TestReceptionExecute_WhenNotRegistered(t *testing.T) {
	mock := &mockDepartment{}
	r := &Reception{}
	r.setNext(mock)
	p := &Patient{name: "test"}
	r.execute(p)
	if !p.registrationDone {
		t.Error("registrationDone not set")
	}
	if !mock.called {
		t.Error("next not called")
	}
	if mock.p != p {
		t.Error("wrong patient passed")
	}
}

func TestReceptionExecute_WhenAlreadyRegistered(t *testing.T) {
	mock := &mockDepartment{}
	r := &Reception{}
	r.setNext(mock)
	p := &Patient{name: "test", registrationDone: true}
	r.execute(p)
	if !mock.called {
		t.Error("next not called even when registered")
	}
}
