package challenge_test

import (
	"io/ioutil"
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

func Test_6_BreakRepeatingKeyXor(t *testing.T) {
	base64Bytes, err := ioutil.ReadFile("6.txt")
	if err != nil {
		t.Error(err)
	}

	inputBytes, err := challenge.Base64ToBytes(string(base64Bytes))
	if err != nil {
		t.Error(err)
	}

	KEYSIZE := challenge.FindProbableKeysize(inputBytes)
	t.Logf("Probable KEYSIZE %d", KEYSIZE)

	// make a block that is the first byte of every block, and a block
	// that is the second byte of every block, and so on.
	parts := [][]byte{}

	for i := 0; i < KEYSIZE; i++ {
		parts = append(parts, []byte{})
	}

	for i, byt := range inputBytes {
		parts[i%KEYSIZE] = append(parts[i%KEYSIZE], byt)
	}

	// Solve each block as if it was single-character XOR.
	key := []byte{}

	for _, part := range parts {
		singleByteKey, _, findErr := challenge.FindSingleByteXorCipherKey(part)
		if findErr != nil {
			t.Error(findErr)
		}

		key = append(key, singleByteKey)
	}

	t.Logf("Probable key `%s`", string(key))

	plaintext, err := challenge.RepeatingKeyXorCipher(inputBytes, key)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(plaintext))
}
