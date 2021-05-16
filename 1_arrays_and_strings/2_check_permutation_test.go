package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Check Permutation: Given two string. Check a string is a permutation of the other.

//Assume 128-character alphabet.
// Time complexity O(n)
// Space complexity O(1)

func CheckPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var charCount [128]int
	for _, c := range a {
		charCount[c] += 1
	}

	for _, c := range b {
		charCount[c] -= 1
		if charCount[c] < 0 {
			return false
		}
	}

	return true
}

func TestCheckPermutation(t *testing.T) {
	testcases := []struct {
		a    string
		b    string
		want bool
	}{
		{
			a:    "abc",
			b:    "bac",
			want: true,
		},
		{
			a:    "abc",
			b:    "abd",
			want: false,
		},
	}

	for _, tc := range testcases {
		got := CheckPermutation(tc.a, tc.b)
		assert.Equalf(t, tc.want, got, "failed CheckPermutation %+v: want: %+v, got: %+v", tc, tc.want, got)
	}
}
