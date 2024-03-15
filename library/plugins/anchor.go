package plugins

import (
	"marceo/library/classes"
	"regexp"
	"strings"
)

var anchorRegex *regexp.Regexp = regexp.MustCompile(`(\[)` + `(.+?)` + `(\]\(.+?\))`)
var anchorHandler func(s, c, e string) string = func(s, c, e string) string {
	url, _ := strings.CutPrefix(e, `](`)
	url, _ = strings.CutSuffix(url, ")")
	return "<a class='markdown anchor' href=\"" + url + "\">" + c + "</a>"
}

var Anchor = classes.NewPlugin("anchor", *anchorRegex, anchorHandler, classes.PluginOptions{})
