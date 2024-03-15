package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var italicBoldRegex *regexp.Regexp = regexp.MustCompile(`(\*\*\*)(.+?)(\*\*\*)`)
var italicBoldHandler func(s, c, e string) string = func(s, c, e string) string { return "<i ><strong>" + c + "</strong></i>" }

var ItalicBold = classes.NewPlugin("itb", *italicBoldRegex, italicBoldHandler, classes.PluginOptions{})
