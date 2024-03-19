package classes

import "sort"

type Pack struct {
	plugins []Plugin
}

func (p Pack) Parse(text string) string {
	newText := "\n" + text + "\n"

	for i := 0; i < len(p.plugins); i++ {
		newText = p.plugins[i].SinalizeText(newText)
	}

	for x := len(p.plugins) - 1; x >= 0; x-- {
		newText = p.plugins[x].ReplaceText(newText)
	}

	return newText
}

func NewPack(pls ...Plugin) Pack {
	sort.Slice(pls, func(i, j int) bool {
		return pls[i].PluginOptions.HideContent
	})
	return Pack{plugins: pls}
}
