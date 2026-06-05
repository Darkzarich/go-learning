package sum

import "testing"

func TestSumPositive(t *testing.T) {
	sum, err := Sum(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestSumNegative(t *testing.T) {
	_, err := Sum(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Sum(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Sum(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestSumZero(t *testing.T) {
	_, err := Sum(0, 0)
	if err == nil {
		t.Error("both args zero - expected error not be nil")
	}
}
