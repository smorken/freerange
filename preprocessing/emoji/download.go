package emoji

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// TwemojiURL gets a formatted Twitter emoji download link
func TwemojiURL(code string) string {
	code = strings.Replace(code, " ", "_", -1)
	code = strings.ToLower(code)
	result := fmt.Sprintf("https://twemoji.maxcdn.com/2/svg/%v.svg", code)
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
