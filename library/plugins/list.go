package plugins

import (
	"marceo/library/classes"
	"regexp"
	"strings"
)

var listRegex *regexp.Regexp = regexp.MustCompile(`(\n)((- .+?\n)+)()`)
var listHandler func(s, c, e string) string = func(s, c, e string) string {

	splited := strings.Split(c, "\n-")

	acum := "<ul class='markdown list'>"
	for _, v := range splited {
		value, _ := strings.CutPrefix(v, "-")
		acum += "<li>" + value + "</li>"
	}
	acum += "</ul>"
	return acum
}

var List = classes.NewPlugin("list", *listRegex, listHandler, classes.PluginOptions{})
