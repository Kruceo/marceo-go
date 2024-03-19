package main

import (
	"log"
	"os"

	"github.com/kruceo/marceo-go/library/packs"
)

const version string = "0.4.2"

func main() {
	var inputText string = ""
	var outputPath string = ""
	for index, v := range os.Args {
		if v == "-i" {
			if len(os.Args) > index+1 {
				argInput := os.Args[index+1]
				file, err := os.ReadFile(argInput)

				if err != nil {
					log.Fatal(err)
				}

				inputText = string(file)
			} else {
				log.Fatal("No input")
			}
		}

		if v == "-o" {
			if len(os.Args) > index+1 {
				argInput := os.Args[index+1]
				outputPath = argInput
			} else {
				log.Fatal("No input")
			}

		}
		if v == "-v" {
			println("v" + version)
			break
		}
		if v == "--help" || v == "-h" {
			println("-i $path$         ", "Read the input path file and parse to html.")
			println("-o $path$         ", "Use the input path to write the results in a file.")
			println("-v                ", "Show version.")
			println("--help -h         ", "Show this message.")
		}
	}

	var parsedText = packs.DefaultPack.Parse(inputText)

	if outputPath != "" {
		os.WriteFile(outputPath, []byte(parsedText), 0644)
	} else if inputText != "" {
		println(parsedText)
	}
}
