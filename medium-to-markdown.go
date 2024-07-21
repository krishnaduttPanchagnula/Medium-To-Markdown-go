package mtm

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// MediumToMarkdown converts a Medium article to Markdown and saves it to a file.
func MediumToMarkdown(URL string, filename string) error {
	converter := md.NewConverter("", true, nil)

	filebytes := CurlFromURL(URL)

	markdowntext, err := converter.ConvertString(string(filebytes))
	if err != nil {
		log.Fatal(err)
	}
	cleanedmarkdown := markDownClean(markdowntext)
	err = saveInFile([]byte(cleanedmarkdown), filename)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}
