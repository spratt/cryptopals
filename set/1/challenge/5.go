package challenge

import (
	"bytes"
	"fmt"
)

// RepeatingKeyXorCipher takes a message []byte and a key []byte of
// any length, creates a []byte with the same length as the message by
// repeating (or truncating) the bytes in key, and XORs this []byte
// against message and returns the result.
func RepeatingKeyXorCipher(message, key []byte) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	for buf.Len() < len(message) {
		buf.Write(key)
	}
	buf.Truncate(len(message))

	keyBytes := buf.Bytes()

	return Xor(message, keyBytes)
}

// BytesToHexString takes []byte and returns a string representing the
// given bytes in hex.
func BytesToHexString(byts []byte) string {
	ret := ""

	for _, byt := range byts {
		ret += fmt.Sprintf("%02x", byt)
	}

	return ret
}
