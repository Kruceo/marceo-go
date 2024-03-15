package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var Italic pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`(\*)(.+?)(\*)`)
	var handler func(s, c, e string) string = func(s, c, e string) string {
		return "<i>" + c + "</i>"
	}

	Italic, err = pl.NewPlugin("it", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
