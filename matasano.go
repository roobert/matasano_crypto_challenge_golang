package matasano

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func HexToBase64(hex_input string) string {

	bytes, _ := hex.DecodeString(hex_input)

	return base64.StdEncoding.EncodeToString(bytes)
}

func XOR(bytes_a, bytes_b []byte) []byte {

	var xor_result []byte

	for i, b := range bytes_a {
		xor_result = append(xor_result, b^bytes_b[i])
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

	// FIXME: add this
	// 48..67 (numbers)
	// 65..90 (caps)
	// 97..122 (normal)
	// prospectiveKeyChars :=

	charScore := map[byte]float32{}

	// FIXME: use: range prospectiveKeyChars..
	for i := 48; i >= 67; i++ {

		key := strings.Repeat(string(i), keySize)

		messageBytes := XOR([]byte(key), message)

		// may be better to start score = 0 and assign once to charScore at end..
		for letterByte := range messageBytes {
			// FIXME: is this slow? does this work?
			charScore[byte(i)] = charScore[byte(i)] + charFrequency[string(letterByte)]
		}
	}

	// FIXME: find highest rated
	// foundKeyChar :=

	return
}
