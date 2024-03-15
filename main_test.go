package main

import (
	"marceo/library/plugins"
	"testing"
)

func TestVas(t *testing.T) {
	// var pack = classes.NewPack(plugins.Table)
	pl := plugins.Table
	res := pl.SinalizeText("\n|header|header|\n|----|----|\n|value|value|\n\n\n\n|header2|header2|\n|------|---|\n|2|2|\n\n\n")
	res = pl.ReplaceText(res)

	println(res)
}
