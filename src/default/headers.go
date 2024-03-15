package defaults

import (
	"log"
	pl "project/src"
	"regexp"
)

var header1Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(# )(.+)(\n)`)
var header2Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(## )(.+)(\n)`)
var header3Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(### )(.+)(\n)`)
var header4Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(#### )(.+)(\n)`)
var header5Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(##### )(.+)(\n)`)
var header6Regx *regexp.Regexp = regexp.MustCompile(`(?:\n)(###### )(.+)(\n)`)

var Header1 pl.Plugin
var Header2 pl.Plugin
var Header3 pl.Plugin
var Header4 pl.Plugin
var Header5 pl.Plugin
var Header6 pl.Plugin
var err error

func init() {
	Header1, err = pl.NewPlugin("h1", *header1Regx, func(s, c, e string) string { return "<h1 class='markdown header1'>" + c + "</h1>" }, pl.PluginOptions{})
	Header2, err = pl.NewPlugin("h2", *header2Regx, func(s, c, e string) string { return "<h2 class='markdown header2'>" + c + "</h2>" }, pl.PluginOptions{})
	Header3, err = pl.NewPlugin("h3", *header3Regx, func(s, c, e string) string { return "<h3 class='markdown header3'>" + c + "</h3>" }, pl.PluginOptions{})
	Header4, err = pl.NewPlugin("h4", *header4Regx, func(s, c, e string) string { return "<h4 class='markdown header4'>" + c + "</h4>" }, pl.PluginOptions{})
	Header5, err = pl.NewPlugin("h5", *header5Regx, func(s, c, e string) string { return "<h5 class='markdown header5'>" + c + "</h5>" }, pl.PluginOptions{})
	Header6, err = pl.NewPlugin("h6", *header6Regx, func(s, c, e string) string { return "<h6 class='markdown header6'>" + c + "</h6>" }, pl.PluginOptions{})

	if err != nil {
		log.Fatal(err)

		// Trate o erro adequadamente aqui, se necessário
	}
}
