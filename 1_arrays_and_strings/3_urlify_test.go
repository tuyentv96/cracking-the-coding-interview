package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//URLIfy: Write a method to replace all spaces in string with '%20'.
//Your are given the 'true' length of string.
//EXAMPLE
//Input: "Mr John Smith    ", 13
//Output: "Mr%20John%20Smith"

//Assume 128-character alphabet.
// Time complexity O(n)
// Space complexity O(1)

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func URLIfy(str string, trueLength int) string {
	index := len(str) - 1
	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str = replaceAtIndex(str, '0', index)
			str = replaceAtIndex(str, '2', index-1)
			str = replaceAtIndex(str, '%', index-2)
			index -= 3
		} else {
			str = replaceAtIndex(str, rune(str[i]), index)
			index--
		}
	}

	return str
}

func TestURLIfy(t *testing.T) {
	testcases := []struct {
		str        string
		trueLength int
		want       string
	}{
		{
			str:        "Mr John Smith    ",
			trueLength: 13,
			want:       "Mr%20John%20Smith",
		},
	}

	for _, tc := range testcases {
		got := URLIfy(tc.str, tc.trueLength)
		assert.Equalf(t, tc.want, got, "failed URLIfy %+v: want: %+v, got: %+v", tc, tc.want, got)
	}
}
