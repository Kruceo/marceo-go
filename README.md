# Marceo Go

[![Build](https://github.com/Kruceo/marceo-go/actions/workflows/build.yml/badge.svg)](https://github.com/Kruceo/marceo-go/actions/workflows/build.yml)
[![GitHub issues](https://img.shields.io/github/issues/kruceo/marceo-go)](https://github.com/kruceo/marceo-go/issues)
[![License](https://img.shields.io/github/license/kruceo/marceo-go)](https://github.com/kruceo/marceo-go/blob/master/LICENSE)
[![Latest Version](https://img.shields.io/github/v/release/kruceo/marceo-go)](https://github.com/kruceo/marceo/releases)

## Description

Marceo-Go is a rewrite from Marceo, a markdown-html parser, but built to `js`. 

- Focused on speed.
- Can be used directly in CLI
- Can be used as dependency.

## Building from sources

```bash
git clone https://github.com/Kruceo/marceo-go.git
cd marceo-go
go build
```

## Basic CLI Usage

```bash 
./marceo -i ./my/text.md -o ./my/output.html
```

## Using in own projects 

### Install

```bash
go get github.com/kruceo/marceo-go
```

### Basic Parse

```go
import "github.com/kruceo/marceo-go/library/packs"

func main(){
    myMarkdown := "# This is a **Title**"

    // The default pack of plugins by Marceo
    pack := packs.DefaultPack

    result := pack.Parse(myMarkdown)

    println(result)
}
```

## Creating my own pack

With Marceo-go you will be able to create a pack with selected (or with your owns) plugins.

```go
import (
	"github.com/kruceo/marceo-go/library/classes"
	"github.com/kruceo/marceo-go/library/plugins"
)

func main() {
	// Added only Headers 1 `#` and Anchors `[text](url)`
	myPack := classes.NewPack(plugins.Header1, plugins.Anchor)

	res1 := myPack.Parse("# My header 1 that will be parsed.")
	res2 := myPack.Parse("## My header 2 that not will be parsed.")
	res3 := myPack.Parse("[My anchor](https://kruceo.com)")

	println(res1)
	println(res2)
	println(res3)
}
```

## Creating a Plugin
To create plugins you will especially need to known a bit of regex expressions.

```go
// This will capture the text between ">" and "<".
var Regex *regexp.Regexp = regexp.MustCompile(`(>)(.+?)(<)`)

// "start", "content" and "end" args.
func Replacer(s, c, e string) string {
	//process your strings.
	return "<myOwnTag>" + c + "</myOwnTag>"
}

// This changes some behaviors on parse
var pluginOptions classes.PluginOptions = classes.PluginOptions{}

// Finally create the plugin
var myPlugin classes.Plugin = classes.NewPlugin("MyPlugin", *Regex, Replacer, pluginOptions)
```

### Testing the plugin

```go
func main() {
	myPack := classes.NewPack(myPlugin)

	res := myPack.Parse("# >This is a test<")

	println(res)
}
```


## Benchmarks

### Benchmark from 26/03/2023 - v0.6.4

The benchmark was run with the `-p` option for `parallelism`.

|Total Words|Markdown Words        |Execution Time |
|-----------|----------------------|---------------|
|70         |50                    |4<sub>ms</sub> |
|734        |500                   |8<sub>ms</sub> |
|6470       |5000                  |57<sub>ms</sub>|


### Testing Machine

|OS|KERNEL|CPU|MEM|STO|
|----|----|----|----|---|
|Arch Linux|Linux 6.7.4-arch1-1|Intel® Core™ i3-1005G1|8,0 GiB|NVME 1200 <sub>mb/s</sub> R&W|