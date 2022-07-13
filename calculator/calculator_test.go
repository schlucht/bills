package calculator_test

import (
	"testing"

	"github.com/schlucht/bills/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name             string
		left, right, sum int
	}{
		{"Positiv Integer", 25, 25, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := calculator.Add(test.left, test.right)
			assert.Equal(t, test.sum, sum)
		})
	}
}
