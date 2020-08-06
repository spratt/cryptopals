package challenge_test

import (
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

const key = "ICE"

func Test_5_ExpectedOutput(t *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26" +
		"226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b2028" +
		"3165286326302e27282f"

	actualBytes, err := challenge.RepeatingKeyXorCipher([]byte(input), []byte(key))
	if err != nil {
		t.Error(err)
	}

	actual := challenge.BytesToHexString(actualBytes)
	if actual != expected {
		t.Errorf("input `%s` key %s expected %s actual %s",
			input, key, expected, actual,
		)
	}
}
