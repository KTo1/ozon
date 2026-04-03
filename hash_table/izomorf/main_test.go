package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_izomorf(t *testing.T) {
	testCases := map[string]struct {
		str1     string
		str2     string
		expected bool
	}{
		"test1": {
			str1:     "egg",
			str2:     "add",
			expected: true,
		},
		"test2": {
			str1:     "foo",
			str2:     "bar",
			expected: false,
		},
		"test3": {
			str1:     "paper",
			str2:     "title",
			expected: true,
		},
		"test4": {
			str1:     "ab",
			str2:     "aa",
			expected: false,
		},
		"test5": {
			str1:     "badc",
			str2:     "baba",
			expected: false,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, v.expected, izomorf(v.str1, v.str2))
		})
	}
}
