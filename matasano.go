package matasano

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

func XORFindSingleCharKey(message []byte) (byte, charData) {
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
	charScore := map[byte]charData{}

	// 0-9
	for b := 48; b != 67; b++ {
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

	// iterate through possible keys and calculate score for each key
	keySize := len(message)

	fmt.Printf("char scores: ")

	for b := range charScore {
		key := strings.Repeat(string(b), keySize)

		messageBytes := XOR([]byte(key), message)

		score := float32(0)

		for _, letterByte := range messageBytes {
			score = score + charFrequency[string(letterByte)]
		}

		charScore[b] = charData{score, messageBytes, b}

		fmt.Printf("[%s: %v] ", string(b), score)
	}

	fmt.Printf("\n")

	// find highest value in map
	highestScore := float32(0)
	var foundKeyChar byte

	for b, data := range charScore {
		if data.score > highestScore {
			highestScore = data.score
			foundKeyChar = b
		}
	}

	fmt.Printf("suspected key: %s, %s\n", string(foundKeyChar), charScore[foundKeyChar])

	return foundKeyChar, charScore[foundKeyChar]
}

func DetectSingleCharacterXOR(fileName string) []byte {

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
		hexMessage := scanner.Text()

		fmt.Printf("\ndecoding: %s\n", hexMessage)

		decodedMessage, _ := hex.DecodeString(hexMessage)

		_, charData := XORFindSingleCharKey(decodedMessage)

		charScore[hexMessage] = charData
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// find message with highest score overall score
	highestScore := float32(0)
	var highestScoringMessage string

	for message, data := range charScore {
		if data.score > highestScore {
			highestScore = data.score
			highestScoringMessage = message
		}
	}

	foundMessage := charScore[highestScoringMessage].decodedMessage

	return foundMessage
}
