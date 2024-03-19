package main

import (
	"log"
	"os"
	"testing"

	"github.com/kruceo/marceo-go/library/packs"
)

func TestVas(t *testing.T) {
	var pack = packs.DefaultPack
	file, err := os.ReadFile("./examples/big.md")
	if err != nil {
		log.Fatal(err)
	}
	println(pack.Parse(string(file)))
}
