package plugins

import (
	"github.com/kruceo/marceo-go/library/classes"
	"regexp"
	"strings"
)

var headerTestRegex = regexp.MustCompile(`^[-]+$`)
var tableRegex *regexp.Regexp = regexp.MustCompile(`(\n\|)((?s).+?)(\|\n[^\|])`)
var tableHandler func(s, c, e string) string = func(s, c, e string) string {
	// println("#" + c + "#")

	lines := strings.Split(c, "\n")
	headerInner := [][]string{}
	bodyInner := [][]string{}
	isHeader := true
	for _, v := range lines {
		v, _ = strings.CutSuffix(v, "|")
		v, _ = strings.CutPrefix(v, "|")
		cells := strings.Split(v, "|")
		lineOutput := []string{}
		for _, cellV := range cells {
			if headerTestRegex.MatchString(cellV) {
				isHeader = false
				continue
			}
			lineOutput = append(lineOutput, cellV)
		}
		if isHeader {
			headerInner = append(headerInner, lineOutput)
		} else {
			bodyInner = append(bodyInner, lineOutput)
		}
	}

	headerAcum := ""
	bodyAcum := ""

	for _, v := range headerInner {
		headerAcum += "<tr>"

		for _, t := range v {
			headerAcum += "<th>" + t + "</th>"
		}

		headerAcum += "</tr>"
	}
	for _, v := range bodyInner {
		bodyAcum += "<tr>"

		for _, t := range v {
			bodyAcum += "<td>" + t + "</td>"
		}

		bodyAcum += "</tr>\n"
	}

	return "<table><thead>" + headerAcum + "</thead>\n<tbody>" + bodyAcum + "</tbody></table>"
}

var Table = classes.NewPlugin("table", *tableRegex, tableHandler, classes.PluginOptions{})
