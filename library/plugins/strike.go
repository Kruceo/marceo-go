package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var strikeRegex *regexp.Regexp = regexp.MustCompile(`(~~)(.+?)(~~)`)
var strikeHandler func(s, c, e string) string = func(_, c, _ string) string {
	return "<s>" + c + "</s>"
}

var Strike = classes.NewPlugin("strike", *strikeRegex, strikeHandler, classes.PluginOptions{})
