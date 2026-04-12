package main

import (
	"log"
	"os"

	"github.com/ok-yyyy/go-identicon"
)

func main() {
	input := "sample"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	data, err := identicon.EncodePNG(input)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("identicon.png", data, 0o0644); err != nil {
		log.Fatal(err)
	}
}
