package packs

import (
	"marceo/library/classes"
	"marceo/library/plugins"
)

var DefaultPack = classes.NewPack(
	plugins.CodeBlock,
	plugins.CodeLine,
	plugins.Header6,
	plugins.Header5,
	plugins.Header4,
	plugins.Header3,
	plugins.Header2,
	plugins.Header1,
	plugins.ItalicBold,
	plugins.Bold,
	plugins.Italic,
	plugins.Italic2,
	plugins.Anchor,
	plugins.List,
	plugins.Table,
)
