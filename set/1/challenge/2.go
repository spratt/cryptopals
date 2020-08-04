package challenge

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrUnequalLengths is the error returned by Xor if the left and
// right values don't have equal length.
var ErrUnequalLengths = errors.New("unequal lengths")

// Xor takes a left and right []byte and returns the resulting []byte
// after XORing the input values together.
func Xor(left, right []byte) ([]byte, error) {
	if len(left) != len(right) {
		return nil, fmt.Errorf(
			"%w: len(left) is %d but len(right) is %d",
			ErrUnequalLengths, len(left), len(right),
		)
	}

	buf := bytes.NewBuffer([]byte{})

	for i := 0; i < len(left); i++ {
		buf.Write([]byte{left[i] ^ right[i]})
	}

	return buf.Bytes(), nil
}
