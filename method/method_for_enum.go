package main

import "fmt"

// DeliveryState — status of delivery.
type DeliveryState string

// Enum for DeliveryState
const (
	// message sent
	DeliveryStatePending DeliveryState = "pending"
	// message received
	DeliveryStateAck DeliveryState = "acknowledged"
	// message processed successfully
	DeliveryStateProcessed DeliveryState = "processed"
	// message canceled
	DeliveryStateCanceled DeliveryState = "canceled"
)

// IsValid checks if the current state is valid
func (s DeliveryState) IsValid() bool {
	switch s {
	case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
		return true
	default:
		return false
	}
}

// String return string representation of the current state.
func (s DeliveryState) String() string {
	return string(s)
}

func HandleMessageDeliveryStatus(status DeliveryState) error {
	if !status.IsValid() {
		return fmt.Errorf("invalid status: %s", status)
	}

	return nil
}

func main() {
	// DeliveryState is a string so we can assign it directly and the
	var currentState DeliveryState = "pending"

	if currentState.IsValid() {
		fmt.Println("Message is sent")
	}

	fmt.Println(currentState.String())

	// assign enum value and call a method

	currentState2 := DeliveryStateAck

	if currentState2.IsValid() {
		fmt.Println("Message is received")
	}

	fmt.Println(currentState2.String())

	// status validation

	if err := HandleMessageDeliveryStatus(DeliveryState("invalid")); err != nil {
		fmt.Println(err)
	}
}
