package packs

import (
	"github.com/kruceo/marceo-go/library/classes"
	"github.com/kruceo/marceo-go/library/plugins"
)

var DefaultPack = classes.NewPack(
	plugins.CodeLine,
	plugins.CodeBlock,
	plugins.Image,
	plugins.List,
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
	plugins.Table,
	plugins.Strike,
	plugins.NumList,
)
