
# Medium to Markdown Converter

A Go program that converts a Medium article to Markdown and saves it to a file.

## Installation

```
go get github.com/krishnaduttPanchagnula/medium-to-markdown-go
```

## Usage

```go
package main

import (
	"log"

	mtm "github.com/krishnaduttPanchagnula/medium-to-markdown-go"
)

func main() {
	err := mtm.MediumToMarkdown("https://medium.com/@krishnaduttpanchagnula/from-scratch-to-brew-creating-personalized-formulae-using-tfblueprintgen-in-homebrew-b5e745b6551d", "test.md")
	if err != nil {
		log.Fatal(err)
	}
}
```

## CLI
 