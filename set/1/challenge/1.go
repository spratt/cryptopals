// Package challenge contains the solutions for cryptopals set 1 challenge 1
package challenge

import (
	"bytes"
	"encoding/base64"
	"strconv"
)

const (
	// HexBase is the number base to use when using hexadecimal digits
	// encoded as a string.
	HexBase = 16

	// ByteBitSize is the number of bits in a byte.
	ByteBitSize = 8
)

// HexStringToBytes takes a string containing a series of hexadecimal
// digits and returns the equivalent []byte.
func HexStringToBytes(hex string) ([]byte, error) {
	// If the length of the string is odd, add a leading zero
	if len(hex)%2 != 0 {
		hex = "0" + hex
	}

	buf := bytes.NewBuffer([]byte{})

	for i := 0; i < len(hex)/2; i++ {
		byteStr := hex[2*i : (2*i)+2]

		curByte, err := strconv.ParseUint(byteStr, HexBase, ByteBitSize)
		if err != nil {
			return nil, err
		}

		buf.Write([]byte{byte(curByte)})
	}

	return buf.Bytes(), nil
}

// BytesToBase64 takes []byte and returns the given []byte encoded in
// a string as base64.
func BytesToBase64(byts []byte) string {
	return base64.StdEncoding.EncodeToString(byts)
}
