package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stack(t *testing.T) {
	testCases := map[string]struct {
		testString string
		expected   bool
	}{
		"()": {
			testString: "()",
			expected:   true},
		"((": {
			testString: "((",
			expected:   false,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, v.expected, isValid(v.testString))
		})
	}
}
