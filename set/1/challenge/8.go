package challenge

import (
	"math"
	"sort"
)

// KeyLength is the assumed length of a key.
const KeyLength = 16

// EqualKeys returns true if the given []byte are equal.
func EqualKeys(firstKey, secondKey []byte) bool {
	if len(firstKey) != len(secondKey) {
		return false
	}

	for i, byt := range firstKey {
		if byt != secondKey[i] {
			return false
		}
	}

	return true
}

// NextKey takes key []byte which is a key of length KeyLength and
// returns the next key by incrementing the 0th byte and making any
// necessary carries.  The maximum key is [255, 255, ..., 255], and if
// NextKey is called on this, then it returns [0, 0, ..., 0] and
// false.
func NextKey(key []byte) ([]byte, bool) {
	if len(key) != KeyLength {
		return key, false
	}

	for i := 0; i < KeyLength; i++ {
		if key[i] < MaxByte {
			key[i]++
			return key, true
		}

		key[i] = 0
	}

	return key, false
}

// FrequencySignature takes a map[byte]float64 containing a map of
// bytes to their relative frequencies in a text represented as
// float64 where 0 means the charcter doesn't occur, and 1 means that
// every character in the text is this character, and returns a list
// of frequencies in descending order which can be used to match
// against character frequencies in a specific language to tell if the
// original text was written in that language.
func FrequencySignature(charFreq map[byte]float64) []float64 {
	sig := []float64{}

	for _, freq := range charFreq {
		sig = append(sig, freq)
	}

	// sort into descending order
	sort.Sort(sort.Reverse(sort.Float64Slice(sig)))

	return sig
}

// CompareFrequencySignatures returns the difference between two
// frequency signatures.
func CompareFrequencySignatures(freq1, freq2 []float64) float64 {
	if len(freq1) > len(freq2) {
		freq1, freq2 = freq2, freq1
	}

	var diff float64

	// can assume freq2 is longer
	for i, freq := range freq1 {
		diff += math.Abs(freq - freq2[i])
	}

	for i := len(freq1); i < len(freq2); i++ {
		diff += freq2[i]
	}

	return diff
}

// ComputeCharacterFrequencies takes an input []byte and returns
// map[byte]float64 which represents the relative frequencies of each
// byte in the input []byte as a float64 where 0 represents the
// character doesn't occur at all, and 1 represents that every
// character in the text is this character.
func ComputeCharacterFrequencies(input []byte) map[byte]float64 {
	freq := make(map[byte]float64)

	for _, byt := range input {
		freq[byt] += 1 / float64(len(input))
	}

	return freq
}

// CrackAES128ECB takes []byte which is English text encrypted using
// AES-128 in EBC mode, and returns a score of how likely the text is
// to be English, the most probable key, and the most probable
// plaintext.
func CrackAES128ECB(input []byte) (float64, []byte, string, error) {
	var (
		bestScore     float64 = math.Inf(-1)
		bestKey       []byte
		bestPlaintext string

		key   = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		valid = true
	)

	for valid {
		plaintext, err := DecryptAes128Ecb(input, key)
		if err != nil {
			return bestScore, bestKey, bestPlaintext, err
		}

		score := Score(plaintext)

		if score > bestScore {
			bestScore = score
			bestKey = key
			bestPlaintext = string(plaintext)
		}

		key, valid = NextKey(key)
	}

	return bestScore, bestKey, bestPlaintext, nil
}
