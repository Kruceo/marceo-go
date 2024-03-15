package defaults

import (
	"log"
	pl "project/src"
	"regexp"
	"strings"
)

var List pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`(\n)((- .+?\n)+)()`)
	var handler func(s, c, e string) string = func(s, c, e string) string {

		splited := strings.Split(c, "\n-")

		acum := "<ul class='markdown list'>"
		for _, v := range splited {
			value, _ := strings.CutPrefix(v, "-")
			acum += "<li>" + value + "</li>"
		}
		acum += "</ul>"
		return acum
	}

	List, err = pl.NewPlugin("list", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
