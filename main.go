// main.go

package main

import (
	"fmt"
	"marceo/library/packs"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File not found")
		return
	}
	str, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Error in file read")
	}
	fmt.Println(packs.DefaultPack.Parse(string(str)))
}
