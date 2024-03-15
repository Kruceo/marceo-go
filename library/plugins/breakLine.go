package plugins

import (
	"marceo/library/classes"

	"regexp"
)

var breakLineRegex *regexp.Regexp = regexp.MustCompile(`()(\n)()`)
var breakLineHandler func(s, c, e string) string = func(s, c, e string) string { return "<br/>" }
var BreakLine = classes.NewPlugin("br", *breakLineRegex, breakLineHandler, classes.PluginOptions{})

// Trate o erro adequadamente aqui, se necessário
