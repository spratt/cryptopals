package challenge

import (
	"math"
	"strings"
)

// SingleByteXorCipher takes a message []byte and a key byte, creates
// a []byte with every byte equal to key, and XORs this []byte against
// message and returns the result.
func SingleByteXorCipher(message []byte, key byte) ([]byte, error) {
	return RepeatingKeyXorCipher(message, []byte{key})
}

// EnglishByteFrequency maps each byte value to a float64 value
// representing how frequently it occurs in English.  These values are
// taken from http://www.fitaly.com/board/domper3/posts/136.html
var EnglishByteFrequency = map[byte]float64{
	9:   0.0057,
	32:  17.1662,
	33:  0.0072,
	34:  0.2442,
	35:  0.0179,
	36:  0.0561,
	37:  0.0160,
	38:  0.0226,
	39:  0.2447,
	40:  0.2178,
	41:  0.2233,
	42:  0.0628,
	43:  0.0215,
	44:  0.7384,
	45:  1.3734,
	46:  1.5124,
	47:  0.1549,
	48:  0.5516,
	49:  0.4594,
	50:  0.3322,
	51:  0.1847,
	52:  0.1348,
	53:  0.1663,
	54:  0.1153,
	55:  0.1030,
	56:  0.1054,
	57:  0.1024,
	58:  0.4354,
	59:  0.1214,
	60:  0.1225,
	61:  0.0227,
	62:  0.1242,
	63:  0.1474,
	64:  0.0073,
	65:  0.3132,
	66:  0.2163,
	67:  0.3906,
	68:  0.3151,
	69:  0.2673,
	70:  0.1416,
	71:  0.1876,
	72:  0.2321,
	73:  0.3211,
	74:  0.1726,
	75:  0.0687,
	76:  0.1884,
	77:  0.3529,
	78:  0.2085,
	79:  0.1842,
	80:  0.2614,
	81:  0.0316,
	82:  0.2519,
	83:  0.4003,
	84:  0.3322,
	85:  0.0814,
	86:  0.0892,
	87:  0.2527,
	88:  0.0343,
	89:  0.0304,
	90:  0.0076,
	91:  0.0086,
	92:  0.0016,
	93:  0.0088,
	94:  0.0003,
	95:  0.1159,
	96:  0.0009,
	97:  5.1880,
	98:  1.0195,
	99:  2.1129,
	100: 2.5071,
	101: 8.5771,
	102: 1.3725,
	103: 1.5597,
	104: 2.7444,
	105: 4.9019,
	106: 0.0867,
	107: 0.6753,
	108: 3.1750,
	109: 1.6437,
	110: 4.9701,
	111: 5.7701,
	112: 1.5482,
	113: 0.0747,
	114: 4.2586,
	115: 4.3686,
	116: 6.3700,
	117: 2.0999,
	118: 0.8462,
	119: 1.3034,
	120: 0.1950,
	121: 1.1330,
	122: 0.0596,
	123: 0.0026,
	124: 0.0007,
	125: 0.0026,
	126: 0.0003,
	149: 0.6410,
	183: 0.0010,
}

// EnglishByteValue gives each ascii byte a value according to its
// frequency in the English language.
func EnglishByteValue(byt byte) float64 {
	if val, ok := EnglishByteFrequency[byt]; ok {
		return val
	}

	return 0
}

// EnglishCharFrequency maps each character value to a float64 value
// representing how frequently it occurs in English.  These values are
// taken from
// https://web.archive.org/web/20040603075055/http://www.data-compression.com/english.html
var EnglishCharFrequency = map[string]float64{
	"a": 0.0651738,
	"b": 0.0124248,
	"c": 0.0217339,
	"d": 0.0349835,
	"e": 0.1041442,
	"f": 0.0197881,
	"g": 0.0158610,
	"h": 0.0492888,
	"i": 0.0558094,
	"j": 0.0009033,
	"k": 0.0050529,
	"l": 0.0331490,
	"m": 0.0202124,
	"n": 0.0564513,
	"o": 0.0596302,
	"p": 0.0137645,
	"q": 0.0008606,
	"r": 0.0497563,
	"s": 0.0515760,
	"t": 0.0729357,
	"u": 0.0225134,
	"v": 0.0082903,
	"w": 0.0171272,
	"x": 0.0013692,
	"y": 0.0145984,
	"z": 0.0007836,
	" ": 0.1918182,
}

// EnglishCharValue gives each character a value according to its
// frequency in the English language.
func EnglishCharValue(byt byte) float64 {
	ch := strings.ToLower(string([]byte{byt}))
	if val, ok := EnglishCharFrequency[ch]; ok {
		return val
	}

	return 0
}

// CharacterwiseScore takes a message []byte and returns a decimal
// value representing how many of the bytes in message are within the
// usual English ASCII characters.
func CharacterwiseScore(message []byte) float64 {
	var ret float64

	for _, byt := range message {
		ret += EnglishCharValue(byt) / float64(len(message))
	}

	return ret
}

// MapDifference returns the numerical difference between two
// map[string]float64.
func MapDifference(firstMap, secondMap map[string]float64) float64 {
	var diff float64
	for key, val := range firstMap {
		diff += math.Abs(val - secondMap[key])
	}

	for key, val := range secondMap {
		if _, ok := firstMap[key]; !ok {
			diff += val
		}
	}

	return diff
}

// Score takes a message []byte and returns a decimal value
// representing how many of the bytes in message are within the usual
// English ASCII characters.
func Score(message []byte) float64 {
	freq := make(map[string]float64)

	for _, byt := range message {
		ch := strings.ToLower(string([]byte{byt}))
		freq[ch] += 1.0 / float64(len(message))
	}

	var penalties float64

	for _, byt := range message {
		if _, ok := EnglishByteFrequency[byt]; !ok {
			penalties -= 1.0 / float64(len(message))
		}
	}

	return 1 - MapDifference(freq, EnglishCharFrequency)
}

const (
	// MinByte is the minimum value byte.
	MinByte byte = 0
	// MaxByte is the maximum value byte.
	MaxByte byte = 255
)

// FindSingleByteXorCipherKey takes []byte that are encrypted using a
// single byte XOR cipher and returns the most likely key given the
// assumption that the plaintext is in English.
func FindSingleByteXorCipherKey(inputBytes []byte) (byte, float64, error) {
	var (
		bestScore float64 = math.Inf(-1)
		bestKey   byte
	)

	for key := MinByte; key < MaxByte; key++ {
		res, err := SingleByteXorCipher(inputBytes, key)
		if err != nil {
			return bestKey, bestScore, err
		}

		score := Score(res)

		if score > bestScore {
			bestScore = score
			bestKey = key
		}
	}

	return bestKey, bestScore, nil
}
