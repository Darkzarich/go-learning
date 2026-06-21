package main

import (
	"errors"
	"testing"
)

type Mock struct {
	response Response
	err      error
}

func NewAPIClientMock() *Mock {
	return &Mock{}
}

func (m *Mock) SetResponse(resp Response, err error) {
	m.response = resp
	m.err = err
}

func (m *Mock) GetData(query string) (Response, error) {
	return m.response, m.err
}

func TestHandleRequest(t *testing.T) {
	mockError := errors.New("something went wrong")

	tests := []struct {
		name      string
		mockedRes Response
		mockedErr error
	}{
		{"returns expected response (200)", Response{Text: "something1", StatusCode: 200}, nil},
		{"returns expected response (404)", Response{Text: "something2", StatusCode: 404}, nil},
		{"returns expected error", Response{}, mockError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewAPIClientMock()
			mock.SetResponse(tt.mockedRes, tt.mockedErr)

			res, err := handleRequest(mock, "")

			if res != tt.mockedRes {
				t.Errorf("handleRequest() ret1 = %v, want %v", res, tt.mockedRes)
			}

			if !errors.Is(err, tt.mockedErr) {
				t.Errorf("handleRequest() ret2 = %v, want %v", err, tt.mockedErr)
			}
		})
	}
}
