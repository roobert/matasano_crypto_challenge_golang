package matasano

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

func HexToBase64(hexInput string) string {
	bytes, _ := hex.DecodeString(hexInput)

	return base64.StdEncoding.EncodeToString(bytes)
}

func XOR(bytes_a, bytes_b []byte) []byte {
	xor_result := make([]byte, len(bytes_a))

	for i, b := range bytes_a {
		xor_result[i] = b ^ bytes_b[i]
	}

	return xor_result
}

type charData struct {
	score          float32
	decodedMessage []byte
	key            byte
}

func XORFindSingleCharKey(message []byte) charData {

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

	// struct to contain data about each char and the likelyhood that
	// each char is the actual key
	charScore := map[byte]charData{}

	// 0-9
	for b := 48; b != 57; b++ {
		charScore[byte(b)] = charData{}
	}

	// caps
	for b := 65; b != 90; b++ {
		charScore[byte(b)] = charData{}
	}

	// a-z
	for b := 97; b != 122; b++ {
		charScore[byte(b)] = charData{}
	}

	// iterate through possible keys and calculate a score for each key
	for charKey := range charScore {

		// key length matches message length
		key := strings.Repeat(string(charKey), len(message))

		messageXOR := XOR([]byte(key), message)

		score := float32(0)

		for _, char := range messageXOR {
			score = score + charFrequency[string(char)]
		}

		// store result for each key
		charScore[charKey] = charData{score, messageXOR, charKey}
	}

	// find the key with associated decoded message that has highest score
	highestScore := float32(0)
	var foundChar charData

	for _, data := range charScore {
		if data.score > highestScore {
			highestScore = data.score
			foundChar = data
		}
	}

	return foundChar
}

func DetectSingleCharacterXOR(fileName string) charData {

	// open file
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// for each line, find most likely key char
	charScore := map[string]charData{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		decodedHexMessage, _ := hex.DecodeString(scanner.Text())

		charData := XORFindSingleCharKey(decodedHexMessage)

		charScore[scanner.Text()] = charData
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// find message with associated decoded message that has the highest score
	highestScore := float32(0)
	var foundChar charData

	for _, data := range charScore {
		if data.score > highestScore {
			highestScore = data.score
			foundChar = data
		}
	}

	return foundChar
}

func RepeatingXOREncrypt(key, message []byte) []byte {

	// create a repeating key that's the same length as the message
	repeatKey := make([]byte, len(message))
	ki := 0

	for i := 0; i != len(message); i++ {
		repeatKey[i] = key[ki]

		if ki == 2 {
			ki = 0
		} else {
			ki++
		}
	}

	return XOR(repeatKey, message)
}
