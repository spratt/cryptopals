package challenge

import (
	"bytes"
)

// SingleByteXorCipher takes a message []byte and a key byte, creates
// a []byte with every byte equal to key, and XORs this []byte against
// message and returns the result.
func SingleByteXorCipher(message []byte, key byte) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	for i := 0; i < len(message); i++ {
		buf.Write([]byte{key})
	}

	keyBytes := buf.Bytes()

	return Xor(message, keyBytes)
}

const englishChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz '\n"

func isEnglishChar(byt byte) bool {
	for _, char := range englishChars {
		if byt == byte(char) {
			return true
		}
	}

	return false
}

// Score takes a message []byte and returns a decimal value
// representing how many of the bytes in message are within the usual
// English ASCII characters.
func Score(message []byte) float64 {
	var ret float64

	for _, byt := range message {
		if isEnglishChar(byt) {
			ret += 1.0 / float64(len(message))
		}
	}

	return ret
}
