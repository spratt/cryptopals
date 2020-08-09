package challenge

import (
	"crypto/aes"
)

// DecryptAes128Ecb takes ciphertext, key []byte and decrypts the
// ciphertext using the key and AES-128 in ECB-mode.
func DecryptAes128Ecb(ciphertext, key []byte) ([]byte, error) {
	decrypted := make([]byte, len(ciphertext))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return decrypted, err
	}

	for start := 0; start < len(ciphertext); start += aes.BlockSize {
		end := start + aes.BlockSize
		cipher.Decrypt(decrypted[start:end], ciphertext[start:end])
	}

	return decrypted, nil
}
