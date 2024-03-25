package plugins

import (
	"regexp"
	"strings"

	"github.com/kruceo/marceo-go/library/classes"
)

var imageRegex *regexp.Regexp = regexp.MustCompile(`(!\[)` + `(.+?)` + `(\]\(.+?\))`)
var imageHandler func(s, c, e string) string = func(s, c, e string) string {
	url, _ := strings.CutPrefix(e, `](`)
	url, _ = strings.CutSuffix(url, ")")
	return "<image src=\"" + url + "\" title=\"" + c + "\"/>"
}

var Image = classes.NewPlugin("image", *imageRegex, imageHandler, classes.PluginOptions{HideContent: true})
