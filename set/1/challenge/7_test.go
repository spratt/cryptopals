package challenge_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_7_DecryptAes128Ecb(t *testing.T) {
	const key = "YELLOW SUBMARINE"

	base64Bytes, err := ioutil.ReadFile("7.txt")
	if err != nil {
		t.Error(err)
	}

	trimmed := strings.Trim(string(base64Bytes), " \n")

	inputBytes, err := challenge.Base64ToBytes(trimmed)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := challenge.DecryptAes128Ecb(inputBytes, []byte(key))
	if err != nil {
		t.Error(err)
	}

	t.Log(string(plaintext))
}
