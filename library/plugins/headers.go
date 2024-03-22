package plugins

import (
	"regexp"

	"github.com/kruceo/marceo-go/library/classes"
)

var header1Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(# )(.+)(\n)`)
var header2Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(## )(.+)(\n)`)
var header3Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(### )(.+)(\n)`)
var header4Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(#### )(.+)(\n)`)
var header5Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(##### )(.+)(\n)`)
var header6Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(###### )(.+)(\n)`)

var Header1 = classes.NewPlugin("h1", *header1Regx, func(s, c, e string) string { return "<h1>" + c + "</h1>" }, classes.PluginOptions{})
var Header2 = classes.NewPlugin("h2", *header2Regx, func(s, c, e string) string { return "<h2>" + c + "</h2>" }, classes.PluginOptions{})
var Header3 = classes.NewPlugin("h3", *header3Regx, func(s, c, e string) string { return "<h3>" + c + "</h3>" }, classes.PluginOptions{})
var Header4 = classes.NewPlugin("h4", *header4Regx, func(s, c, e string) string { return "<h4>" + c + "</h4>" }, classes.PluginOptions{})
var Header5 = classes.NewPlugin("h5", *header5Regx, func(s, c, e string) string { return "<h5>" + c + "</h5>" }, classes.PluginOptions{})
var Header6 = classes.NewPlugin("h6", *header6Regx, func(s, c, e string) string { return "<h6>" + c + "</h6>" }, classes.PluginOptions{})
