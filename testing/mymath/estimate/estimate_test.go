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
}
