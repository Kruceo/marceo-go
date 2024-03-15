package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var italicRegex *regexp.Regexp = regexp.MustCompile(`(\*)(.+?)(\*)`)
var italicHandler func(s, c, e string) string = func(s, c, e string) string {
	return "<i>" + c + "</i>"
}

var Italic = classes.NewPlugin("it", *italicRegex, italicHandler, classes.PluginOptions{})
