package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var CodeLine pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile("(`)" + `(.+?)` + "(`)")
	var handler func(s, c, e string) string = func(s, c, e string) string {
		return "<code>" + c + "</code>"
	}

	CodeLine, err = pl.NewPlugin("CodeLine", *regex, handler, pl.PluginOptions{HideContent: true})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
