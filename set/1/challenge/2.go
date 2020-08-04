package challenge

import (
	"bytes"
	"errors"
	"fmt"
)

var ErrUnequalLengths = errors.New("unequal lengths")

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
