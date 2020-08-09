package challenge_test

import (
	"io/ioutil"
	"math"
	"strings"
	"testing"

	"github.com/spratt/cryptopals/set/1/challenge"
)

func Test_8_EqualKeys(t *testing.T) {
	testCases := []struct {
		FirstKey      []byte
		SecondKey     []byte
		ExpectedEqual bool
	}{
		{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
		{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}

	for _, testCase := range testCases {
		actualEqual := challenge.EqualKeys(testCase.FirstKey, testCase.SecondKey)

		if actualEqual != testCase.ExpectedEqual {
			t.Errorf("key %v expectedEqual %t actualEqual %t",
				key, testCase.ExpectedEqual, actualEqual,
			)
		}
	}
}

func Test_8_NextKey(t *testing.T) {
	testCases := []struct {
		Key             []byte
		ExpectedNextKey []byte
		ExpectedValid   bool
	}{
		{[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
		{
			[]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
	}

	for _, testCase := range testCases {
		actualNextKey, actualValid := challenge.NextKey(testCase.Key)

		if !challenge.EqualKeys(actualNextKey, testCase.ExpectedNextKey) {
			t.Errorf("key %v expectedNextKey %v actualNextKey %v",
				key, testCase.ExpectedNextKey, actualNextKey,
			)
		}

		if actualValid != testCase.ExpectedValid {
			t.Errorf("key %v expectedValid %t actualValid %t",
				key, testCase.ExpectedValid, actualValid,
			)
		}
	}
}

var expectedSig = []float64{
	0.171662,
	0.085771,
	0.063700,
	0.057701,
	0.051880,
	0.049701,
	0.049019,
	0.043686,
	0.042586,
	0.031750,
	0.027444,
	0.025071,
	0.021129,
	0.020999,
	0.016437,
	0.015597,
	0.015482,
	0.015124,
	0.013734,
	0.013725,
	0.013034,
	0.011330,
	0.010195,
	0.008462,
	0.007384,
	0.006753,
	0.006410,
	0.005516,
	0.004594,
	0.004354,
	0.004003,
	0.003906,
	0.003529,
	0.003322,
	0.003322,
	0.003211,
	0.003151,
	0.003132,
	0.002673,
	0.002614,
	0.002527,
	0.002519,
	0.002447,
	0.002442,
	0.002321,
	0.002233,
	0.002178,
	0.002163,
	0.002085,
	0.001950,
	0.001884,
	0.001876,
	0.001847,
	0.001842,
	0.001726,
	0.001663,
	0.001549,
	0.001474,
	0.001416,
	0.001348,
	0.001242,
	0.001225,
	0.001214,
	0.001159,
	0.001153,
	0.001054,
	0.001030,
	0.001024,
	0.000892,
	0.000867,
	0.000814,
	0.000747,
	0.000687,
	0.000628,
	0.000596,
	0.000561,
	0.000343,
	0.000316,
	0.000304,
	0.000227,
	0.000226,
	0.000215,
	0.000179,
	0.000160,
	0.000088,
	0.000086,
	0.000076,
	0.000073,
	0.000072,
	0.000057,
	0.000026,
	0.000026,
	0.000016,
	0.000010,
	0.000009,
	0.000007,
	0.000003,
	0.000003,
}

func Test_8_FrequencySignature(t *testing.T) {
	actualSig := challenge.FrequencySignature(challenge.EnglishByteFrequency)

	if challenge.CompareFrequencySignatures(actualSig, expectedSig) > 0 {
		t.Logf("Unexpected difference between expected signature %f and actual %f",
			expectedSig, actualSig,
		)
	}
}

var testText = `I'm back and I'm ringin' the bell 
        A rockin' on the mike while the fly girls yell 
        In ecstasy in the back of me 
        Well that's my DJ Deshay cuttin' all them Z's 
        Hittin' hard and the girlies goin' crazy 
        Vanilla's on the mike, man I'm not lazy. 
        
        I'm lettin' my drug kick in 
        It controls my mouth and I begin 
        To just let it flow, let my concepts go 
        My posse's to the side yellin', Go Vanilla Go! 
        
        Smooth 'cause that's the way I will be 
        And if you don't give a damn, then 
        Why you starin' at me 
        So get off 'cause I control the stage 
        There's no dissin' allowed 
        I'm in my own phase 
        The girlies sa y they love me and that is ok 
        And I can dance better than any kid n' play 
        
        Stage 2 -- Yea the one ya' wanna listen to 
        It's off my head so let the beat play through 
        So I can funk it up and make it sound good 
        1-2-3 Yo -- Knock on some wood 
        For good luck, I like my rhymes atrocious 
        Supercalafragilisticexpialidocious 
        I'm an effect and that you can bet 
        I can take a fly girl and make her wet. 
        
        I'm like Samson -- Samson to Delilah 
        There's no denyin', You can try to hang 
        But you'll keep tryin' to get my style 
        Over and over, practice makes perfect 
        But not if you're a loafer. 
        
        You'll get nowhere, no place, no time, no girls 
        Soon -- Oh my God, homebody, you probably eat 
        Spaghetti with a spoon! Come on and say it! 
        
        VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino 
        Intoxicating so you stagger like a wino 
        So punks stop trying and girl stop cryin' 
        Vanilla Ice is sellin' and you people are buyin' 
        'Cause why the freaks are jockin' like Crazy Glue 
        Movin' and groovin' trying to sing along 
        All through the ghetto groovin' this here song 
        Now you're amazed by the VIP posse. 
        
        Steppin' so hard like a German Nazi 
        Startled by the bases hittin' ground 
        There's no trippin' on mine, I'm just gettin' down 
        Sparkamatic, I'm hangin' tight like a fanatic 
        You trapped me once and I thought that 
        You might have it 
        So step down and lend me your ear 
        '89 in my time! You, '90 is my year. 
        
        You're weakenin' fast, YO! and I can tell it 
        Your body's gettin' hot, so, so I can smell it 
        So don't be mad and don't be sad 
        'Cause the lyrics belong to ICE, You can call me Dad 
        You're pitchin' a fit, so step back and endure 
        Let the witch doctor, Ice, do the dance to cure 
        So come up close and don't be square 
        You wanna battle me -- Anytime, anywhere 
        
        You thought that I was weak, Boy, you're dead wrong 
        So come on, everybody and sing this song 
        
        Say -- Play that funky music Say, go white boy, go white boy go 
        play that funky music Go white boy, go white boy, go 
        Lay down and boogie and play that funky music till you die. 
        
        Play that funky music Come on, Come on, let me hear 
        Play that funky music white boy you say it, say it 
        Play that funky music A little louder now 
        Play that funky music, white boy Come on, Come on, Come on 
        Play that funky music`

func Test_8_ComputeCharacterFrequencies(t *testing.T) {
	freq := challenge.ComputeCharacterFrequencies([]byte(testText))
	sig := challenge.FrequencySignature(freq)
	englishSig := challenge.FrequencySignature(challenge.EnglishByteFrequency)
	sigDiff := challenge.CompareFrequencySignatures(sig, englishSig)

	if sigDiff > 0.4 {
		t.Logf("Unexpected signature difference %f between expected signature %f and actual %f",
			sigDiff, englishSig, sig,
		)
	}
}

func Test_8_DetectAes128Ecb(t *testing.T) {
	hexBytes, err := ioutil.ReadFile("8.txt")
	if err != nil {
		t.Error(err)
	}

	lines := strings.Split(string(hexBytes), "\n")

	englishSig := challenge.FrequencySignature(challenge.EnglishByteFrequency)

	var (
		bestLine  []byte
		bestScore float64 = math.Inf(1)
	)

	for _, line := range lines {
		inputBytes, lineErr := challenge.HexStringToBytes(line)
		if lineErr != nil {
			t.Error(lineErr)
		}

		freq := challenge.ComputeCharacterFrequencies(inputBytes)
		sig := challenge.FrequencySignature(freq)
		score := challenge.CompareFrequencySignatures(sig, englishSig)

		if score < bestScore {
			bestScore = score
			bestLine = inputBytes
		}
	}

	t.Logf("With a score of %f the following line is probably AES-128 EBC encrypted english: %v",
		bestScore, bestLine,
	)

	// This test shows how to brute-force the input, but it would take
	// forever so let's skip it.
	t.SkipNow()

	score, key, plaintext, err := challenge.CrackAES128ECB(bestLine)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Cracked score %f key %v plaintext %s",
		score, key, plaintext,
	)
}

func Test_8_DecryptAes128Ecb(t *testing.T) {
	// This test shows how to brute-force the input, but it would take
	// forever so let's skip it.
	t.SkipNow()

	hexBytes, err := ioutil.ReadFile("8.txt")
	if err != nil {
		t.Error(err)
	}

	lines := strings.Split(string(hexBytes), "\n")

	var (
		bestPlaintext string
		bestScore     float64 = math.Inf(-1)
	)

	for i, line := range lines {
		inputBytes, err := challenge.HexStringToBytes(line)
		if err != nil {
			t.Error(err)
		}

		score, _, plaintext, err := challenge.CrackAES128ECB(inputBytes)
		if err != nil {
			t.Error(err)
		}

		if score > bestScore {
			bestScore = score
			bestPlaintext = plaintext
		}

		if i%50 == 0 {
			t.Logf("Decrypted %d/%d lines, best score so far: %f", i+1, len(lines), bestScore)
		}
	}

	t.Logf("best score %f best plaintext: %s", bestScore, bestPlaintext)
}
