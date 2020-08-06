package challenge

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
