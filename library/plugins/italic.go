package plugins

import (
	"github.com/kruceo/marceo-go/library/classes"
	"regexp"
)

var italicRegex *regexp.Regexp = regexp.MustCompile(`(\*)(.+?)(\*)`)
var italic2Regex *regexp.Regexp = regexp.MustCompile(`(_)(.+?)(_)`)
var italicHandler func(s, c, e string) string = func(s, c, e string) string {
	return "<i>" + c + "</i>"
}

var Italic = classes.NewPlugin("it", *italicRegex, italicHandler, classes.PluginOptions{})
var Italic2 = classes.NewPlugin("it", *italic2Regex, italicHandler, classes.PluginOptions{})
