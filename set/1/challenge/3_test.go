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

	var (
		bestScore float64
		bestKey   byte
		bestRes   []byte
	)

	for key := byte(0); key < byte(255); key++ {
		res, err := challenge.SingleByteXorCipher(inputBytes, key)
		if err != nil {
			t.Error(err)
		}

		score := challenge.Score(res)

		if score > bestScore {
			bestScore = score
			bestKey = key
			bestRes = res
		}
	}

	t.Logf("Score %f key %d string `%s`", bestScore, bestKey, string(bestRes))
}
