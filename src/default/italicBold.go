package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var ItalicBold pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`(\*\*\*)(.+?)(\*\*\*)`)
	var handler func(s, c, e string) string = func(s, c, e string) string { return "<i class='markdown italic'><strong>" + c + "</strong></i>" }

	ItalicBold, err = pl.NewPlugin("itb", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
