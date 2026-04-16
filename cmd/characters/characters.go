// Package main provides the characters CLI.
package main

import (
	"flag"
	"github.com/arran4/sentencestats/pkg/cli"
)

var (
	outputFile = flag.String("o", "out.png", "Output PNG file name")
)

func main() {
	flag.Parse()
	cli.Characters(*outputFile)
}
