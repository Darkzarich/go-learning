package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetData(query string) (Response, error) {
	args := m.Called(query)
	return args.Get(0).(Response), args.Error(1)
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
			mockClient := new(Mock)
			mockClient.On("GetData", "test query").Return(tt.mockedRes, tt.mockedErr)

			res, err := handleRequest(mockClient, "test query")

			assert.Equal(t, tt.mockedRes, res)
			assert.ErrorIs(t, tt.mockedErr, err)
			mockClient.AssertExpectations(t)
		})
	}
}
