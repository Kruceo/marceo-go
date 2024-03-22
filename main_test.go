package main

import (
	"log"
	"os"
	"testing"

	"github.com/kruceo/marceo-go/library/packs"
)

func TestVas(t *testing.T) {
	mk, err := os.ReadFile("ztest.md")
	if err != nil {
		log.Fatal(err)
	}
	text := string(mk)

	text = packs.DefaultPack.ParalelParse(text)
	
	println(text)
}
