package pl

import (
	"errors"
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

func NewPlugin(name string, regex regexp.Regexp, elHandler func(s string, c string, e string) string, pluginOptions PluginOptions) (Plugin, error) {

	if regex.NumSubexp() < 3 {
		return Plugin{}, errors.New("A plugin need 3 regexp groups. Pl: " + name + ".")
	}

	return Plugin{
		Name:          name,
		mem:           make(map[string][]string),
		Regex:         regex,
		Handler:       elHandler,
		PluginOptions: pluginOptions,
	}, nil
}

func (p *Plugin) SinalizeText(text string) string {

	ntxt := strings.Clone(text)

	var catched [][]string = p.Regex.FindAllStringSubmatch(ntxt, -1)
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

		ntxt = regexp.MustCompile(regexp.QuoteMeta(full)).ReplaceAllLiteralString(ntxt, tokenizedContent)
	}
	return ntxt
}

func (p *Plugin) ReplaceText(text string) string {
	ntxt := strings.Clone(text)
	for key, val := range p.mem {
		regExp := regexp.MustCompile("(?s)(" + regexp.QuoteMeta(key) + ")(.*?)(" + regexp.QuoteMeta(key) + ")")
		keyMatch := regExp.FindStringSubmatch(ntxt)
		if len(keyMatch) < 4 {
			continue
		}

		if p.PluginOptions.HideContent {
			el := p.Handler(val[1], val[2], val[3])
			ntxt = strings.ReplaceAll(ntxt, keyMatch[0], el)

		} else {
			el := p.Handler(val[1], keyMatch[2], val[3])
			ntxt = strings.ReplaceAll(ntxt, keyMatch[0], el)
		}

	}
	return ntxt

}
