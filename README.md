
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
	err := mtm.MediumToMarkdown("<YOUR_URL>", "test.md")
	if err != nil {
		log.Fatal(err)
	}
}
```

## CLI
If you are not interested in writing go code and just want to convert the files using our CLI tool
```bash
mtm --url <YOUR_URL> --filename test.md
```
