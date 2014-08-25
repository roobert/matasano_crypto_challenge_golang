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
	bytesA, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	bytesB, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	hexExpectedXORResult := "746865206b696420646f6e277420706c6179"

	XORResult := XOR(bytesA, bytesB)

	if hex.EncodeToString(XORResult) != hexExpectedXORResult {
		t.Error("xor result does not match expected xor result")
	}
}

func TestXORFindSingleCharKey(t *testing.T) {
	inputMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedMessage := "Cooking MC's like a pound of bacon"
	decodedInputMessage, _ := hex.DecodeString(inputMessage)

	charData := XORFindSingleCharKey(decodedInputMessage)

	if string(charData.decodedMessage) != expectedMessage {
		t.Error("decrypted message doesn't match expected message")
	}
}

func TestDetectSingleCharacterXOR(t *testing.T) {
	expectedMessage := "Now that the party is jumping\n"
	charData := DetectSingleCharacterXOR("data/4/gistfile1.txt")

	if string(charData.decodedMessage) != expectedMessage {
		t.Error("found message doesn't match expected message")
	}
}

func TestRepeatingXOREncrypt(t *testing.T) {
	message := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	expectedEncryptedMessage := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	encryptedMessage := RepeatingXOREncrypt([]byte(key), []byte(message))

	if hex.EncodeToString(encryptedMessage) != expectedEncryptedMessage {
		t.Error("encrypted message doesn't match expected encrypted message")
	}
}

// 1-6
func TestHammingDistance(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"

	if hammingDistance([]byte(a), []byte(b)) != 37 {
		t.Error("hamming distance did not match expected hamming distance")
	}
}

func TestCalculateKeySize(t *testing.T) {
	t.Skip()
}

func TestBreakRepeatingXOR(t *testing.T) {

}
