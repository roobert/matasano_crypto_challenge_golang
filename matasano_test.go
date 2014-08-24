package matasano

import (
	"encoding/hex"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64Output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	base64Output := HexToBase64(hexInput)

	if base64Output != expectedBase64Output {
		t.Error("base64 output does not match expected output!")
	}
}

func TestXOR(t *testing.T) {
	hexInputA := "1c0111001f010100061a024b53535009181c"
	hexInputB := "686974207468652062756c6c277320657965"

	hexExpectedXORResult := "746865206b696420646f6e277420706c6179"

	bytesA, _ := hex.DecodeString(hexInputA)
	bytesB, _ := hex.DecodeString(hexInputB)

	XORResult := XOR(bytesA, bytesB)
	hexXORResult := hex.EncodeToString(XORResult)

	if hexXORResult != hexExpectedXORResult {
		t.Error("xor result does not match expected xor result")
	}
}

func TestXORFindSingleCharKey(t *testing.T) {
	inputMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	expectedMessage := "Cooking MC's like a pound of bacon"

	rawMessage, _ := hex.DecodeString(inputMessage)

	_, charData := XORFindSingleCharKey(rawMessage)

	if string(charData.decodedMessage) != expectedMessage {
		t.Error("decrypted message doesn't match expected message")
	}
}

func TestDetectSingleCharacterXOR(t *testing.T) {

	expectedMessage := "Now that the party is jumping\n"

	foundMessage := DetectSingleCharacterXOR("data/4/gistfile1.txt")

	if string(foundMessage) != expectedMessage {
		t.Error("found message doesn't match expected message")
	}
}
