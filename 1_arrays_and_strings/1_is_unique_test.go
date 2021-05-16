package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
//cannot use additional data structures?

//Assume 128-character alphabet.
// Time complexity O(n)
// Space complexity O(1)

func IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}

	var charSet [128]bool
	for _, c := range str {
		if charSet[c] {
			return false
		}

		charSet[c] = true
	}

	return true
}

// Use bit vector
// Use lowercase letters from a to z
// Time complexity O(n)
// Space complexity O(1)

func IsUnique2(str string) bool {
	if len(str) > 128 {
		return false
	}

	var checker int
	for _, c := range str {
		val := c - 'a'
		if (checker & (1 << val)) > 0 {
			return false
		}

		checker |= 1 << val
	}

	return true
}

func TestIsUnique(t *testing.T) {
	testcases := []struct {
		str  string
		want bool
	}{
		{
			str:  "abc",
			want: true,
		},
		{
			str:  "aab",
			want: false,
		},
		{
			str:  "abcdeefgh",
			want: false,
		},
	}

	for _, tc := range testcases {
		got := IsUnique(tc.str)
		assert.Equalf(t, tc.want, got, "failed IsUnique %+v: want: %+v, got: %+v", tc, tc.want, got)
	}

	for _, tc := range testcases {
		got := IsUnique2(tc.str)
		assert.Equalf(t, tc.want, got, "failed IsUnique2 %+v: want: %+v, got: %+v", tc, tc.want, got)
	}
}
