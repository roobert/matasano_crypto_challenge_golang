package matasano

import (
	"encoding/base64"
	"encoding/hex"
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
