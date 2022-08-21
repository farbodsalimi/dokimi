package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "zero",
			input:  "0",
			expect: 0,
		},
		{
			name:   "one,two,three",
			input:  "123",
			expect: 123,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, ToInt(testCase.input))
		})
	}
}
