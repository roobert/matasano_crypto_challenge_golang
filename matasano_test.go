package matasano

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	hex_input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected_base64_output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	base64_output := HexToBase64(hex_input)

	if base64_output != expected_base64_output {
		t.Error("base64_output does not match expected output!")
	}
}

func TestXOR(t *testing.T) {
	hex_input_a := "1c0111001f010100061a024b53535009181c"
	hex_input_b := "686974207468652062756c6c277320657965"

	hex_expected_xor_result := "746865206b696420646f6e277420706c6179"

	bytes_a, _ := hex.DecodeString(hex_input_a)
	bytes_b, _ := hex.DecodeString(hex_input_b)

	xor_result := XOR(bytes_a, bytes_b)
	hex_xor_result := hex.EncodeToString(xor_result)

	if hex_xor_result != hex_expected_xor_result {
		t.Error("xor result does not match expected xor result")
	}
}

func TestXORFindSingleCharKey(t *testing.T) {
	inputMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	expectedMessage := "Cooking MC's like a pound of bacon"

	rawMessage, _ := hex.DecodeString(inputMessage)

	foundKeyChar := XORFindSingleCharKey(rawMessage)

	probableKey := strings.Repeat(string(foundKeyChar), len(inputMessage)/2)

	decryptedMessage := XOR([]byte(probableKey), rawMessage)

	if string(decryptedMessage) != expectedMessage {
		t.Error("decrypted message doesn't match expected message")
	}
}
