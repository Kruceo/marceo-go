package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var boldRegex *regexp.Regexp = regexp.MustCompile(`(\*\*)(.+?)(\*\*)`)
var boldHandler func(s, c, e string) string = func(s, c, e string) string { return "<strong>" + c + "</strong>" }

var Bold = classes.NewPlugin("b", *boldRegex, boldHandler, classes.PluginOptions{})
