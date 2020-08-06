package challenge_test

import (
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_6_HammingDistance(t *testing.T) {
	testCases := []struct {
		Left     string
		Right    string
		Expected int
	}{
		{"", "", 0},
		{"abc", "", 24},
		{"", "abc", 24},
		{"abc", "abc", 0},
		{"abc", "bbc", 2},
		{"abc", "acc", 1},
		{"abc", "abd", 3},
		{"abc", "bcd", 6},
		{"abc", "abcdef", 24},
		{"abc", "bcddef", 30},
		{"this is a test", "wokka wokka!!!", 37},
	}

	for _, testCase := range testCases {
		actual := challenge.HammingDistance(testCase.Left, testCase.Right)
		if actual != testCase.Expected {
			t.Errorf("left `%s` right `%s` expected %d actual %d",
				testCase.Left, testCase.Right, testCase.Expected, actual,
			)
		}
	}
}
