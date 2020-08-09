package challenge

import (
	"encoding/base64"
	"sort"
)

// HammingDistance takes two strings and returns the number of bits
// that differ in their byte representation.
func HammingDistance(left, right string) int {
	dist := 0

	if len(left) > len(right) {
		dist += (len(left) - len(right)) * ByteBitSize
		left = left[:len(right)]
	} else if len(right) > len(left) {
		dist += (len(right) - len(left)) * ByteBitSize
		right = right[:len(left)]
	}

	for i := 0; i < len(left); i++ {
		diff := left[i] ^ right[i]
		for shift := 0; shift < 8; shift++ {
			if (diff>>shift)&1 == 1 {
				dist++
			}
		}
	}

	return dist
}

// Base64ToBytes takes a string containing base64-encoded bytes and
// returns those bytes and any decoding error that occurred.
func Base64ToBytes(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// KeysizeDist is a structure containing the combination of a given
// Keysize and a measure of its Distance.
type KeysizeDist struct {
	Keysize int
	Dist    float64
}

// FindProbableKeysize takes a []byte ciphertext encrypted using a
// repeating key XOR cipher and finds the probable keysize.
func FindProbableKeysize(inputBytes []byte) int {
	const (
		minKeysize = 1
		maxKeysize = 60

		topNKeysizes       = 10
		partsToInvestigate = 10
	)

	ndists := []KeysizeDist{}

	// Let KEYSIZE be the guessed length of the key; try values from 2
	// to (say) 40.
	for KEYSIZE := minKeysize; KEYSIZE <= maxKeysize; KEYSIZE++ {
		// For each KEYSIZE, take the first KEYSIZE worth of bytes,
		// and the second KEYSIZE worth of bytes, and find the edit
		// distance between them. Normalize this result by dividing by
		// KEYSIZE.
		firstPart := inputBytes[:KEYSIZE]
		secondPart := inputBytes[KEYSIZE : KEYSIZE*2]
		dist := HammingDistance(string(firstPart), string(secondPart))
		ndist := float64(dist) / float64(KEYSIZE)
		ndists = append(ndists, KeysizeDist{KEYSIZE, ndist})
	}

	// The KEYSIZE with the smallest normalized edit distance is
	// probably the key.
	sort.SliceStable(ndists, func(i, j int) bool {
		return ndists[i].Dist < ndists[j].Dist
	})

	// Take the most likely N KEYSIZEs
	mostLikelyKeysizes := []int{}

	for i := 0; i < topNKeysizes; i++ {
		mostLikelyKeysizes = append(mostLikelyKeysizes, ndists[i].Keysize)
	}

	// For each of these most likely KEYSIZEs, take more KEYSIZE
	// blocks and compute the average hamming distance between them.
	avgNdists := []KeysizeDist{}

	for _, KEYSIZE := range mostLikelyKeysizes {
		parts := [][]byte{}

		for i := 0; i < partsToInvestigate; i++ {
			parts = append(parts, inputBytes[KEYSIZE*i:KEYSIZE*(i+1)])
		}

		ndists := []float64{}

		for i, part := range parts {
			for j := i + 1; j < len(parts); j++ {
				dist := HammingDistance(
					string(part), string(parts[j]),
				)
				ndists = append(ndists, float64(dist)/float64(KEYSIZE))
			}
		}

		var avgNdist float64

		for _, ndist := range ndists {
			avgNdist += ndist / float64(len(ndists))
		}

		avgNdists = append(avgNdists, KeysizeDist{KEYSIZE, avgNdist})
	}

	// The KEYSIZE with the smallest normalized edit distance is
	// probably the key.
	sort.SliceStable(avgNdists, func(i, j int) bool {
		return avgNdists[i].Dist < avgNdists[j].Dist
	})

	return avgNdists[0].Keysize
}
