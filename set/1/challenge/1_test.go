package challenge_test

import (
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_ExpectedOutput(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	hexBytes, err := challenge.HexStringToBytes(input)
	if err != nil {
		t.Error(err)
	}
	actual := challenge.BytesToBase64(hexBytes)
	if actual != expected {
		t.Errorf("input %s expected %s actual %s", input, expected, actual)
	}
}
