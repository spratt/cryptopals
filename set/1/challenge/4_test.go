package challenge_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_4_FindOutput(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("4.txt")
	if err != nil {
		t.Error(err)
	}

	inputLines := strings.Split(string(inputBytes), "\n")

	var (
		bestInput []byte
		bestRes   []byte
		bestScore float64
	)

	for _, input := range inputLines {
		inputBytes, err := challenge.HexStringToBytes(input)
		if err != nil {
			t.Error(err)
		}

		for key := byte(0); key < byte(255); key++ {
			res, err := challenge.SingleByteXorCipher(inputBytes, key)
			if err != nil {
				t.Error(err)
			}

			score := challenge.Score(res)

			if score > bestScore {
				bestScore = score
				bestInput = inputBytes
				bestRes = res
			}
		}
	}

	t.Logf("Score %f string %s input %v", bestScore, string(bestRes), bestInput)
}
