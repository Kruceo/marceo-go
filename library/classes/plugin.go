package classes

import (
	"regexp"
	"strconv"
	"strings"
)

type PluginOptions struct {
	HideContent bool
}

type Plugin struct {
	mem           map[string][]string
	Regex         regexp.Regexp
	Name          string
	Handler       func(start string, content string, end string) string
	PluginOptions PluginOptions
}

func NewPlugin(name string, regex regexp.Regexp, elHandler func(s string, c string, e string) string, pluginOptions PluginOptions) Plugin {
	regGrpNum := regex.NumSubexp()
	if regGrpNum < 3 {
		regStr := regex.String()
		for i := 0; i < 3-regGrpNum; i++ {
			regStr += "()"
		}
		regex = *regexp.MustCompile(regStr)
	}
	return Plugin{
		Name:          name,
		mem:           make(map[string][]string),
		Regex:         regex,
		Handler:       elHandler,
		PluginOptions: pluginOptions,
	}
}

func (p *Plugin) SinalizeText(text string) string {

	var catched [][]string = p.Regex.FindAllStringSubmatch(text, -1)

	for i := 0; i < len(catched); i++ {
		full := catched[i][0]
		start := catched[i][1]
		content := catched[i][2]
		end := catched[i][3]
		token := "@" + p.Name + "-" + strconv.Itoa(i) + "@"
		p.mem[token] = []string{full, start, content, end}

		tokenizedContent := token + content + token

		if p.PluginOptions.HideContent {
			tokenizedContent = token + token
		}

		text = regexp.MustCompile(regexp.QuoteMeta(full)).ReplaceAllLiteralString(text, tokenizedContent)
	}
	return text
}

func (p *Plugin) ReplaceText(text string) string {
	for key, val := range p.mem {
		regExp := regexp.MustCompile("(?s)(" + regexp.QuoteMeta(key) + ")(.*?)(" + regexp.QuoteMeta(key) + ")")
		keyMatch := regExp.FindStringSubmatch(text)
		if len(keyMatch) < 4 {
			continue
		}

		if p.PluginOptions.HideContent {
			el := p.Handler(val[1], val[2], val[3])
			text = strings.ReplaceAll(text, keyMatch[0], el)

		} else {
			el := p.Handler(val[1], keyMatch[2], val[3])
			text = strings.ReplaceAll(text, keyMatch[0], el)
		}

	}
	return text

}
