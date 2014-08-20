package matasano

import "testing"

func TestHexToBase64(t *testing.T) {
	hex_input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected_base64_output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	base64_output := HexToBase64(hex_input)

	if base64_output != expected_base64_output {
		t.Error("base64_output does not match expected output!")
	}
}
