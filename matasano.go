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
