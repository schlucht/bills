package codewars_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/schlucht/bills/codewars"
	"github.com/stretchr/testify/assert"
)

func PrintLine(function string) {
	l := len(function)
	total := 80
	rep := total - l

	fmt.Println()
	fmt.Println(strings.Repeat("=", rep/2), function, strings.Repeat("=", rep/2))
	fmt.Println()
}
func TestReverseString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "first test",
			input:  "The quick brown fox jumps over the lazy dog.",
			output: "ehT kciuq nworb xof spmuj revo eht yzal .god",
		},
		{
			name:   "second test",
			input:  "apple",
			output: "elppa",
		},
		{
			name:   "three test",
			input:  "a b c d",
			output: "a b c d",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := codewars.ReverseString(test.input)
			assert.Equal(t, res, test.output)
		})
	}
}

func TestGetSum(t *testing.T) {
	PrintLine("TestGetSum")
	tests := []struct {
		name            string
		intA, intB, out int
	}{
		{"1.", 0, 1, 1},
		{"2.", 505, 4, 127759},
		{"3.", -321, 123, -44055},
		{"4.", 1, 1, 1},
		{"5.", -505, 4, -127755},
		{"6.", 5, 1, 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := codewars.GetSum(test.intA, test.intB)
			assert.Equal(t, res, test.out)
		})
	}
}

func TestIpsBetween(t *testing.T) {
	PrintLine("TestIpsBetween")
	tests := []struct {
		name string
		a, b string
		out  int
	}{
		{"1.", "10.0.0.0", "10.0.0.50", 50},
		{"2.", "20.0.0.10", "20.0.1.0", 246},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := codewars.IpsBetween(test.a, test.a)
			assert.Equal(t, res, test.out)
		})
	}
}

func TestNextBigger(t *testing.T) {
	PrintLine("NextBigger")
	tests := []struct {
		in, out int
	}{
		{12, 21},
		{513, 531},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.in), func(t *testing.T) {
			res := codewars.NextBigger(test.in)
			assert.Equal(t, res, test.out)
		})
	}
}
