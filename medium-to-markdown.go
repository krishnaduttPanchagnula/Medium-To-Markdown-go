package mtm

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

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

// func CurlfromURL(url string) []byte {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	bodyText, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s\n", bodyText)
// 	return bodyText
// }

// func saveinfile(filebody []byte, filename string) error {
// 	err := os.WriteFile(filename, filebody, os.ModeAppend.Perm())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return nil
// }
