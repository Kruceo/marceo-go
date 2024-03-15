package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var Bold pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`(\*\*)(.+?)(\*\*)`)
	var handler func(s, c, e string) string = func(s, c, e string) string { return "<strong class='markdown bold'>" + c + "</strong>" }

	Bold, err = pl.NewPlugin("b", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
