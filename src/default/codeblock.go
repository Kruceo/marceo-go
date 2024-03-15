package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var CodeBlock pl.Plugin

func init() {
	var quotesStr = regexp.QuoteMeta("```")
	var regex *regexp.Regexp = regexp.MustCompile(`(?s)(\n` + quotesStr + `\w+[ |\n])(.+?)(` + quotesStr + "\n)")
	var handler func(s, c, e string) string = func(s, c, e string) string {
		return "<pre class='markdown codeblock'><code>" + c + "</code></pre>"
	}

	CodeBlock, err = pl.NewPlugin("codeblock", *regex, handler, pl.PluginOptions{HideContent: true})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
