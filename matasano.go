package matasano

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hex_input string) string {

	bytes, _ := hex.DecodeString(hex_input)

	base64_output := base64.StdEncoding.EncodeToString(bytes)

	return base64_output
}

func XOR(bytes_a, bytes_b []byte) []byte {

	var xor_result []byte

	for i, b := range bytes_a {
		xored_byte := b ^ bytes_b[i]

		xor_result = append(xor_result, xored_byte)
	}

	return xor_result
}
