// Package main provides the character-pairs CLI.
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
	cli.CharacterPairs(*outputFile)
}
