package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//One Way: There are three types of edit that can be performed on string, insert, remove and replace
//Given two string. Write a function to check if they are one edit or zero edit
//EXAMPLE
//pale, ple -> true
//pales, pale -> true
//pale, bale -> true
//pale, bae -> false

// Time Complexity O(n)
// Space Complexity O(1)

func oneWayReplace(a, b string) bool {
	foundDif := false
	for i := range a {
		if a[i] != b[i] {
			if foundDif {
				return false
			}

			foundDif = true
		}
	}

	return true
}

// len(a) > len(b)
func oneWayInsert(a, b string) bool {
	i1, i2 := 0, 0
	for i1 < len(a) && i2 < len(b) {
		if a[i1] != b[i2] {
			if i1 != i2 {
				return false
			}

			i1++
		} else {
			i1++
			i2++
		}
	}

	return true
}

func OneWay(a, b string) bool {
	la, lb := len(a), len(b)
	if la == lb {
		return oneWayReplace(a, b)
	} else if la > lb {
		return oneWayInsert(a, b)
	}

	return oneWayInsert(b, a)
}

//pale, ple -> true
//pales, pale -> true
//pale, bale -> true
//pale, bae -> false
func TestOneWay(t *testing.T) {
	testcases := []struct {
		a    string
		b    string
		want bool
	}{
		{
			a:    "pale",
			b:    "ple",
			want: true,
		},
		{
			a:    "pales",
			b:    "pale",
			want: true,
		},
		{
			a:    "pale",
			b:    "bale",
			want: true,
		},
		{
			a:    "pale",
			b:    "bae",
			want: false,
		},
	}

	for _, tc := range testcases {
		got := OneWay(tc.a, tc.b)
		assert.Equalf(t, tc.want, got, "failed OneWay %+v: want: %+v, got: %+v", tc, tc.want, got)
	}
}
