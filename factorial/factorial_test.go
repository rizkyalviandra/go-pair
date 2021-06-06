package factorial_test

import (
	"testing"

	"github.com/rizkyalviandra/go-pair/factorial"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	t.Run("factorial above 0", func(t *testing.T) {
		c := factorial.Factorial(5)
		expect := 120
		assert.Equal(t,expect, c, "expected %d but got %d", expect, c)
	})

	t.Run("factorial is 0", func(t *testing.T) {
		c := factorial.Factorial(0)
		expect := 1
		assert.Equal(t,expect, c, "expected %d but got %d", expect, c)
	})

	t.Run("factorial below 0", func(t *testing.T) {
		c := factorial.Factorial(-5)
		expect := -5
		assert.Equal(t,expect, c, "expected %d but got %d", expect, c)
	})
}