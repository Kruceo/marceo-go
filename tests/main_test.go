package tests

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	// file, err := os.ReadFile("text.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	for i := 0; i < b.N; i++ {
		// Parse(string(file))
	}

}
