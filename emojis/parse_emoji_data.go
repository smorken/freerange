package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Emoji emoji code and name data
type Emoji struct {
	Group       string
	SubGroup    string
	Description string
	Code        string
}

func main() {
	file, err := os.Open("emoji-test.txt")
	check(err)
	defer file.Close()
	var currentGroup = ""
	var currentSubGroup = ""
	descriptionRegEx := regexp.MustCompile("[a-zA-Z ]+$")
	emojiCodeRegEx := regexp.MustCompile("[0-9A-Za-z ]+ +;")
	data := []Emoji{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "# group: ") {
			currentGroup = strings.TrimPrefix(line, "# group: ")
			continue
		} else if strings.HasPrefix(line, "# subgroup: ") {
			currentSubGroup = strings.TrimPrefix(line, "# subgroup: ")
			continue
		} else if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		emojicode := strings.TrimSpace(strings.TrimSuffix(emojiCodeRegEx.FindString(line), ";"))
		description := strings.TrimSpace(descriptionRegEx.FindString(line))

		data = append(data, Emoji{
			Group:       currentGroup,
			SubGroup:    currentSubGroup,
			Description: description,
			Code:        emojicode})
	}

	err = scanner.Err()
	check(err)

	jsonData, err := json.Marshal(data)
	check(err)

	err = ioutil.WriteFile("output.json", jsonData, 777)
}