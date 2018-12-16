package emoji

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ProcessCode converts a code from the specification format to the twitter format
func ProcessCode(code string) string {
	code = strings.Replace(code, " ", "-", -1)
	return strings.ToLower(code)
}

// TwemojiURL gets a formatted Twitter emoji download link
func TwemojiURL(code string, svg bool) string {
	code = ProcessCode(code)
	result := ""
	if svg {
		result = fmt.Sprintf("https://twemoji.maxcdn.com/2/svg/%v.svg", code)
	} else {
		result = fmt.Sprintf("https://twemoji.maxcdn.com/2/72x72/%v.png", code)
	}

	return result
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
