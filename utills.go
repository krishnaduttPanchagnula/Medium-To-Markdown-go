package mtm

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func markDownClean(filebody string) string {
	var shareFound bool
	var result strings.Builder

	lines := strings.Split(filebody, "\n")
	for _, line := range lines {
		if !shareFound {
			if strings.Contains(line, "Share") {
				shareFound = true
			} else {
				continue
			}
		} else {
			result.WriteString(line + "\n")
			if strings.Contains(line, "Follow") {
				break
			}
		}
	}

	return result.String()
}

func CurlFromURL(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	return bodyText
}

func saveInFile(filebody []byte, filename string) error {
	err := os.WriteFile(filename, filebody, os.ModeAppend.Perm())
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
