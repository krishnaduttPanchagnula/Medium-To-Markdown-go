package main

import (
	"flag"
	"fmt"
	"log"

	mtm "github.com/krishnaduttPanchagnula/medium-to-markdown-go"
)

func main() {
	urlFlag := flag.String("url", "", "Medium article URL")
	filenameFlag := flag.String("filename", "", "Destination file name")
	flag.Parse()

	if *urlFlag == "" || *filenameFlag == "" {
		fmt.Println("Please provide a URL and a file name")
		return
	}

	err := mtm.MediumToMarkdown(*urlFlag, *filenameFlag)
	if err != nil {
		log.Fatal(err)
	}
}
