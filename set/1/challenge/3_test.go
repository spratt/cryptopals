package challenge_test

import (
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_3_FindOutput(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	inputBytes, err := challenge.HexStringToBytes(input)
	if err != nil {
		t.Error(err)
	}

	key, score, err := challenge.FindSingleByteXorCipherKey(inputBytes)
	if err != nil {
		t.Error(err)
	}

	res, err := challenge.SingleByteXorCipher(inputBytes, key)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Key %d score %f res `%s`", key, score, res)
}
