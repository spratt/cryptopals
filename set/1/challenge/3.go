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
	9:   0.000057,
	32:  0.171662,
	33:  0.000072,
	34:  0.002442,
	35:  0.000179,
	36:  0.000561,
	37:  0.000160,
	38:  0.000226,
	39:  0.002447,
	40:  0.002178,
	41:  0.002233,
	42:  0.000628,
	43:  0.000215,
	44:  0.007384,
	45:  0.013734,
	46:  0.015124,
	47:  0.001549,
	48:  0.005516,
	49:  0.004594,
	50:  0.003322,
	51:  0.001847,
	52:  0.001348,
	53:  0.001663,
	54:  0.001153,
	55:  0.001030,
	56:  0.001054,
	57:  0.001024,
	58:  0.004354,
	59:  0.001214,
	60:  0.001225,
	61:  0.000227,
	62:  0.001242,
	63:  0.001474,
	64:  0.000073,
	65:  0.003132,
	66:  0.002163,
	67:  0.003906,
	68:  0.003151,
	69:  0.002673,
	70:  0.001416,
	71:  0.001876,
	72:  0.002321,
	73:  0.003211,
	74:  0.001726,
	75:  0.000687,
	76:  0.001884,
	77:  0.003529,
	78:  0.002085,
	79:  0.001842,
	80:  0.002614,
	81:  0.000316,
	82:  0.002519,
	83:  0.004003,
	84:  0.003322,
	85:  0.000814,
	86:  0.000892,
	87:  0.002527,
	88:  0.000343,
	89:  0.000304,
	90:  0.000076,
	91:  0.000086,
	92:  0.000016,
	93:  0.000088,
	94:  0.000003,
	95:  0.001159,
	96:  0.000009,
	97:  0.051880,
	98:  0.010195,
	99:  0.021129,
	100: 0.025071,
	101: 0.085771,
	102: 0.013725,
	103: 0.015597,
	104: 0.027444,
	105: 0.049019,
	106: 0.000867,
	107: 0.006753,
	108: 0.031750,
	109: 0.016437,
	110: 0.049701,
	111: 0.057701,
	112: 0.015482,
	113: 0.000747,
	114: 0.042586,
	115: 0.043686,
	116: 0.063700,
	117: 0.020999,
	118: 0.008462,
	119: 0.013034,
	120: 0.001950,
	121: 0.011330,
	122: 0.000596,
	123: 0.000026,
	124: 0.000007,
	125: 0.000026,
	126: 0.000003,
	149: 0.006410,
	183: 0.000010,
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
