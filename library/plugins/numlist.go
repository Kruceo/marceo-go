package plugins

import (
	"marceo/library/classes"
	"regexp"
	"strings"
)

var spliterRegex *regexp.Regexp = regexp.MustCompile(`\n\d\. `)

var numlistRegex *regexp.Regexp = regexp.MustCompile(`(\n)((( )*?\d\. .+?\n)+)()`)

func numlistHandler(s, c, e string) string {
	acum := "<ol>"
	splited := spliterRegex.Split("\n"+c, -1)[1:]
	for _, v := range splited {
		values := strings.SplitN(v, "\n", 2)

		// if len(values) > 1 && values[1] != nil && numlistRegex.MatchString(values[1]) {
		// println("#------")
		// v = strings.TrimSpace(v)
		// acum += numlistHandler("", strings.TrimPrefix(v, " "), "")
		// continue
		// }

		acum += "<li>" + values[0] + "</li>"

		if len(values) > 1 && numlistRegex.MatchString(values[1]) {
			acum += "<li>"
			newValueSplited := strings.Split(values[1], "\n")
			newValue := ""

			for _, x := range newValueSplited {
				xres, _ := strings.CutPrefix(x, "    ")
				newValue += xres + "\n"
			}
			acum += numlistHandler("", newValue, "")
			acum += "</li>"
		}
	}
	acum += "</ol>"

	return acum
}

var NumList = classes.NewPlugin("numlist", *numlistRegex, numlistHandler, classes.PluginOptions{})
