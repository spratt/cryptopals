package challenge

import (
	"bytes"
	"fmt"
)

func Xor(left, right []byte) ([]byte, error) {
	if len(left) != len(right) {
		return nil, fmt.Errorf(
			"Unequal lengths, len(left) is %d but len(right) is %d",
			len(left), len(right),
		)
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < len(left); i++ {
		buf.Write([]byte{left[i] ^ right[i]})
	}
	return buf.Bytes(), nil
}
