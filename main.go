// main.go

package main

import (
	"fmt"
	"os"
	pl "project/src"
	defaults "project/src/default"
)

func Parse(text string) string {
	newText := "\n" + text + "\n"
	pls := [...]pl.Plugin{
		defaults.CodeBlock,
		defaults.CodeLine,
		defaults.Header6,
		defaults.Header5,
		defaults.Header4,
		defaults.Header3,
		defaults.Header2,
		defaults.Header1,
		defaults.ItalicBold,
		defaults.Bold,
		defaults.Italic,
		defaults.Anchor,
		defaults.List,
	}

	for i := 0; i < len(pls); i++ {
		newText = pls[i].SinalizeText(newText)
	}

	for x := len(pls) - 1; x >= 0; x-- {
		newText = pls[x].ReplaceText(newText)
	}

	return newText
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File not found")
		return
	}
	str, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Error in file read")
	}
	fmt.Println(Parse(string(str)))
}
