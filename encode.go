package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	inputString := os.Args[1]
	encodedString := b64encode(inputString)
	fmt.Println(encodedString)
}

func b64encode(inputString string) string {
	encoder := base64.StdEncoding
	encoded := encoder.EncodeToString([]byte(inputString))
	return encoded
}
