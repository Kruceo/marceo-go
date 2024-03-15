package plugins

import (
	"marceo/library/classes"
	"regexp"
)

var header1Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(# )(.+)(\n)`)
var header2Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(## )(.+)(\n)`)
var header3Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(### )(.+)(\n)`)
var header4Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(#### )(.+)(\n)`)
var header5Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(##### )(.+)(\n)`)
var header6Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(###### )(.+)(\n)`)

var Header1 = classes.NewPlugin("header1", *header1Regx, func(s, c, e string) string { return "<h1 class='markdown header1'>" + c + "</h1>" }, classes.PluginOptions{})
var Header2 = classes.NewPlugin("header2", *header2Regx, func(s, c, e string) string { return "<h2 class='markdown header2'>" + c + "</h2>" }, classes.PluginOptions{})
var Header3 = classes.NewPlugin("header3", *header3Regx, func(s, c, e string) string { return "<h3 class='markdown header3'>" + c + "</h3>" }, classes.PluginOptions{})
var Header4 = classes.NewPlugin("header4", *header4Regx, func(s, c, e string) string { return "<h4 class='markdown header4'>" + c + "</h4>" }, classes.PluginOptions{})
var Header5 = classes.NewPlugin("header5", *header5Regx, func(s, c, e string) string { return "<h5 class='markdown header5'>" + c + "</h5>" }, classes.PluginOptions{})
var Header6 = classes.NewPlugin("header6", *header6Regx, func(s, c, e string) string { return "<h6 class='markdown header6'>" + c + "</h6>" }, classes.PluginOptions{})
