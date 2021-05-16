package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//URLIfy: Given a string, check it is a permutation of a palindrome
//EXAMPLE
//Input: tact coa
//Output: True (permutations: "taco cat'; "atco eta·; etc.)

func GetCharNumber(c rune) int {
	a := 'a'
	z := 'z'
	if c >= a && c <= z {
		return int(c - a)
	}

	return -1
}

// Time Complexity O(n)
// Space Complexity O(1)

func PalindromePermutation(str string) bool {
	charCount := make([]int, GetCharNumber('a')+GetCharNumber('z')+1)
	for _, c := range str {
		x := GetCharNumber(c)
		if x != -1 {
			charCount[x] += 1
		}

	}

	odd := false
	for _, count := range charCount {
		if count%2 > 0 {
			if odd == true {
				return false
			}

			odd = true
		}
	}

	return true
}

// Time Complexity O(n)
// Space Complexity O(1)

func PalindromePermutation2(str string) bool {
	charCount := make([]int, GetCharNumber('a')+GetCharNumber('z')+1)
	countOdd := 0

	for _, c := range str {
		x := GetCharNumber(c)
		if x != -1 {
			charCount[x] += 1
			if charCount[x]%2 > 0 {
				countOdd += 1
			} else {
				countOdd -= 1
			}
		}

	}

	return countOdd <= 1
}

func toggleBitVector(bitVector int, index int) int {
	mask := 1 << index
	if bitVector&mask == 0 {
		bitVector |= mask
	} else {
		bitVector &= ^mask
	}

	return bitVector
}

func createBitVector(str string) int {
	var bitVector int
	for _, c := range str {
		x := GetCharNumber(c)
		if x != -1 {
			bitVector = toggleBitVector(bitVector, x)
		}
	}

	return bitVector
}

// Wre can check the number only has one bit set because if we subtract by 1 then AND it with new number, we should get 0.
// 00010000 - 1 = 00001111
// 00010000 & 00001111 = 0

func checkExactlyOneBitSet(bitVector int) bool {
	return bitVector&(bitVector-1) == 0
}

func PalindromePermutation3(str string) bool {
	bitVector := createBitVector(str)

	return bitVector == 0 || checkExactlyOneBitSet(bitVector)
}

func TestPalindromePermutation(t *testing.T) {
	testcases := []struct {
		str  string
		want bool
	}{
		{
			str:  "tact coa",
			want: true,
		},
		{
			str:  "aabbcd",
			want: false,
		},
	}

	for _, tc := range testcases {
		got := PalindromePermutation(tc.str)
		assert.Equalf(t, tc.want, got, "failed PalindromePermutation %+v: want: %+v, got: %+v", tc, tc.want, got)
	}

	for _, tc := range testcases {
		got := PalindromePermutation2(tc.str)
		assert.Equalf(t, tc.want, got, "failed PalindromePermutation2 %+v: want: %+v, got: %+v", tc, tc.want, got)
	}

	for _, tc := range testcases {
		got := PalindromePermutation3(tc.str)
		assert.Equalf(t, tc.want, got, "failed PalindromePermutation3 %+v: want: %+v, got: %+v", tc, tc.want, got)
	}
}
