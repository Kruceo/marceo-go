package defaults

import (
	"log"
	pl "project/src"
	"regexp"
	"strings"
)

var Anchor pl.Plugin

func init() {

	var regex *regexp.Regexp = regexp.MustCompile(`(\[)` + `(.+?)` + `(\]\(.+?\))`)
	var handler func(s, c, e string) string = func(s, c, e string) string {
		url, _ := strings.CutPrefix(e, `](`)
		url, _ = strings.CutSuffix(url, ")")
		return "<a class='markdown anchor' href=\"" + url + "\">" + c + "</a>"
	}

	Anchor, err = pl.NewPlugin("it", *regex, handler, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}