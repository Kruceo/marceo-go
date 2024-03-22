package classes

import (
	"regexp"
	"sort"
	"strings"
	"sync"
)

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

var lineBlockReg regexp.Regexp = *regexp.MustCompile(`(#)+ .*?`)

func (p Pack) ParalelParse(text string) string {
	parts := []string{}
	partsAdded := 0
	acum := ""
	splited := strings.Split(text, "\n")
	for _, v := range splited {

		if partsAdded > 200 {
			if lineBlockReg.MatchString(v) {

				parts = append(parts, acum)
				partsAdded = 0
				acum = ""

			}
		}

		acum += v + "\n"
		partsAdded++

	}

	partsCopy := make([]string, len(parts))
	copy(partsCopy, parts)

	wg := sync.WaitGroup{}
	wg.Add(len(partsCopy))
	for i := 0; i < len(partsCopy); i++ {
		partsCopy[i] = p.Parse(partsCopy[i])
		wg.Done()
	}
	wg.Wait()

	acum = ""
	for _, v := range partsCopy {
		acum += v
	}
	return acum
}

func NewPack(pls ...Plugin) Pack {
	sort.Slice(pls, func(i, j int) bool {
		return pls[i].PluginOptions.HideContent
	})
	return Pack{plugins: pls}
}
