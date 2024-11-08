package main

import (
	"encoding/base64"
	"syscall/js"
)

func decodeBase64(this js.Value, p []js.Value) interface{} {
	encodedStr := p[0].String()
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return js.ValueOf("Error: Invalid Base64 string.")
	}
	return js.ValueOf(string(decodedBytes))
}

func main() {
	js.Global().Set("decodeBase64", js.FuncOf(decodeBase64))
	select {} // keep Go WASM running
}
