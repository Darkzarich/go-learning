package estimate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEstimateValue(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		result := EstimateValue(5)

		assert.Equal(t, "small", result)
	})

	t.Run("medium", func(t *testing.T) {
		result := EstimateValue(50)

		assert.Equal(t, "medium", result)
	})

	t.Run("big", func(t *testing.T) {
		result := EstimateValue(100)

		assert.Equal(t, "big", result)
	})

	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Cases
		{
			name: "small",
			args: args{
				value: 5,
			},
			want: "small",
		},
		{
			name: "medium",
			args: args{
				value: 50,
			},
			want: "medium",
		},
		{
			name: "big",
			args: args{
				value: 100,
			},
			want: "big",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateValue(tt.args.value); got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
