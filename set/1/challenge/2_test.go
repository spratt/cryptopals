package challenge_test

import (
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_2_ExpectedOutput(t *testing.T) {
	left := "1c0111001f010100061a024b53535009181c"
	right := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	leftBytes, err := challenge.HexStringToBytes(left)
	if err != nil {
		t.Error(err)
	}

	rightBytes, err := challenge.HexStringToBytes(right)
	if err != nil {
		t.Error(err)
	}

	expectedBytes, err := challenge.HexStringToBytes(expected)
	if err != nil {
		t.Error(err)
	}

	actual, err := challenge.Xor(leftBytes, rightBytes)
	if err != nil {
		t.Error(err)
	}

	actualBase64 := challenge.BytesToBase64(actual)
	expectedBase64 := challenge.BytesToBase64(expectedBytes)

	if actualBase64 != expectedBase64 {
		t.Errorf("left %s right %s expected %s actual %s",
			left, right, expectedBase64, actualBase64,
		)
	}
}
