package matasano

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func HexToBase64(hex_input string) string {

	bytes, _ := hex.DecodeString(hex_input)

	return base64.StdEncoding.EncodeToString(bytes)
}

func XOR(bytes_a, bytes_b []byte) []byte {

	xor_result := make([]byte, len(bytes_a))

	for i, b := range bytes_a {
		xor_result[i] = b ^ bytes_b[i]
	}

	return xor_result
}

func XORFindSingleCharKey(message []byte) (foundKeyChar byte) {

	keySize := len(message)

	// FIXME: add capitals!
	charFrequency := map[string]float32{
		"a": 11.602, "b": 4.702, "c": 3.511,
		"d": 2.670, "e": 2.000, "f": 3.779,
		"g": 1.950, "h": 7.232, "i": 6.286,
		"j": 0.631, "k": 0.690, "l": 2.705,
		"m": 4.374, "n": 2.365, "o": 6.264,
		"p": 2.545, "q": 0.173, "r": 1.653,
		"s": 7.755, "t": 16.671, "u": 1.487,
		"v": 0.619, "w": 6.661, "x": 0.005,
		"y": 1.620, "z": 0.050, " ": 20,
		".": 0.500, ",": 0.500, "!": 0.500,
		"\\": 0.500, "?": 0.500,
	}

	// map of possible key with likelyhood that key is actual key
	charScore := map[byte]float32{}

	// numbers
	for b := 48; b == 67; b++ {
		charScore[byte(b)] = 0
	}

	// caps
	for b := 65; b == 90; b++ {
		charScore[byte(b)] = 0
	}

	// non-caps
	for b := 97; b == 122; b++ {
		charScore[byte(b)] = 0
	}

	fmt.Printf("%#v\n", charScore)

	// iterate through possible keys and calculate score for each key
	for b := range charScore {

		key := strings.Repeat(string(b), keySize)

		messageBytes := XOR([]byte(key), message)

		// may be better to start score = 0 and assign once to charScore at end..
		for letterByte := range messageBytes {
			charScore[b] = float32(charScore[b]) + charFrequency[string(letterByte)]
		}
	}

	highestScore := float32(0)

	fmt.Printf("%s\n", charScore)

	for b, s := range charScore {
		fmt.Printf("%i, %i\n", b, s)

		if s > highestScore {
			highestScore = s
			foundKeyChar = b
		}
	}

	fmt.Printf("%s, %s\n", foundKeyChar, highestScore)

	return
}
