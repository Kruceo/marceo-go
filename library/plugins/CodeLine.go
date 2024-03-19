package plugins

import (
	"github.com/kruceo/marceo-go/library/classes"
	"regexp"
)

var codeLineRegex *regexp.Regexp = regexp.MustCompile("(`)" + `(.+?)` + "(`)")
var codeLineHandler func(s, c, e string) string = func(s, c, e string) string {
	return "<code>" + c + "</code>"
}

var CodeLine = classes.NewPlugin("CodeLine", *codeLineRegex, codeLineHandler, classes.PluginOptions{HideContent: true})

// Trate o erro adequadamente aqui, se
