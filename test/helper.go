package main

import (
	"encoding/json"
	"io"
)

func getPayload(container interface{}, input io.Reader) {
	decodedInput := json.NewDecoder(input)
	decodedInput.Decode(&container) // skipped error check
}
