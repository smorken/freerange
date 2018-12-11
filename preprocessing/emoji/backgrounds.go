package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

func loadNames(nameList string) []string {
	namesDat, err := ioutil.ReadFile(nameList)
	check(err)
	names := []string{}
	err = json.Unmarshal(namesDat, &names)
	check(err)
	return names
}

//ProcessBackgroundEmojis reads in name list and processes specified emojis
func ProcessBackgroundEmojis(nameList string, emojidata string, outDir string) {
	var emojis []Emoji
	dat, err := ioutil.ReadFile(emojidata)
	check(err)
	err = json.Unmarshal(dat, &emojis)
	check(err)
	byName := map[string]Emoji{}
	for i := 0; i < len(emojis); i++ {
		byName[emojis[i].Description] = emojis[i]
	}
	for _, name := range loadNames(nameList) {
		match := byName[name]
		DownloadFile(path.Join(outDir, match.Description)+".svg", TwemojiURL(match.Code))
		fmt.Println(match.Code + " " + match.Description + " " + match.Group)
	}

}
