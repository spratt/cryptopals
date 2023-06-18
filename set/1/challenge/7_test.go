package challenge_test

import (
	"os"
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_7_DecryptAes128Ecb(t *testing.T) {
	const key = "YELLOW SUBMARINE"

	base64Bytes, err := os.ReadFile("7.txt")
	if err != nil {
		t.Error(err)
	}

	inputBytes, err := challenge.Base64ToBytes(string(base64Bytes))
	if err != nil {
		t.Error(err)
	}

	plaintext, err := challenge.DecryptAes128Ecb(inputBytes, []byte(key))
	if err != nil {
		t.Error(err)
	}

	t.Log(string(plaintext))
}
