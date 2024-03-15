package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var BreakLine pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`()(\n)()`)
	var handler func(s, c, e string) string = func(s, c, e string) string { return "<br/>" }

	BreakLine, err = pl.NewPlugin("br", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
