package plugins

import (
	"regexp"
	"strings"

	"github.com/kruceo/marceo-go/library/classes"
)

var imageRegex *regexp.Regexp = regexp.MustCompile(`(!\[)` + `(.+?)` + `(\]\(.+?\))`)
var imageHandler func(s, c, e string) string = func(s, c, e string) string {
	rawurl, _ := strings.CutPrefix(e, `](`)
	rawurl, _ = strings.CutSuffix(rawurl, ")")
	splitedUrl := strings.Split(rawurl, " ")
	url := splitedUrl[0]
	alt := c
	if len(splitedUrl) > 1 {
		alt = splitedUrl[1]
	}
	return "<image src=\"" + url + "\" title=\"" + c + "\" alt=\"" + alt + "\"/>"
}

var Image = classes.NewPlugin("image", *imageRegex, imageHandler, classes.PluginOptions{HideContent: true})
