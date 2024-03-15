package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var quotesStr = regexp.QuoteMeta("```")
var codeblockRegex *regexp.Regexp = regexp.MustCompile(`(?s)(\n` + quotesStr + `\w+[ |\n])(.+?)(` + quotesStr + "\n)")
var codeblockHandler func(s, c, e string) string = func(s, c, e string) string {
	return "<pre class='markdown codeblock'><code>" + c + "</code></pre>"
}

var CodeBlock = classes.NewPlugin("codeblock", *codeblockRegex, codeblockHandler, classes.PluginOptions{HideContent: true})

// Trate o erro adequadamente aqui, se necessário
