package main

import (
	"log"
	"marceo/library/packs"
	"os"
	"testing"
)

func TestVas(t *testing.T) {
	var pack = packs.DefaultPack
	file, err := os.ReadFile("./examples/big.md")
	if err != nil {
		log.Fatal(err)
	}
	println(pack.Parse(string(file)))
}
