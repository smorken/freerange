package emoji

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
		DownloadFile(path.Join(outDir, match.Description)+".svg", TwemojiURL(match.Code, true))
		fmt.Println(match.Code + " " + match.Description + " " + match.Group)
	}

}
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

//ProcessEmojis produces a data file for each emoji subgroup for downloading emoji images
func ProcessEmojis(emojidata string, outdir string) {
	var emojis []Emoji
	dat, err := ioutil.ReadFile(emojidata)
	check(err)
	err = json.Unmarshal(dat, &emojis)
	check(err)
	bySubGroup := map[string][]Emoji{}
	for i := 0; i < len(emojis); i++ {
		subgroupname := emojis[i].SubGroup
		if subgroup, ok := bySubGroup[subgroupname]; ok {
			bySubGroup[subgroupname] = append(subgroup, emojis[i])
		} else {
			bySubGroup[subgroupname] = []Emoji{emojis[i]}
		}
	}
	for subgroupname, subgroup := range bySubGroup {

		for i := 0; i < len(subgroup); i++ {
			emojiFileName := ProcessCode(subgroup[i].Code) + ".png"
			pathToEmoji := path.Join("D:", "CODE", "twemoji", "2", "72x72", emojiFileName)
			if _, err := os.Stat(pathToEmoji); os.IsNotExist(err) {
				//file does not exist so dont use it
				continue
			}
			subgroupOutDir := path.Join(outdir, subgroupname)
			fmt.Printf("group[%s] name[%s] url[%s]\n", subgroupname, subgroup[i].Description, TwemojiURL(subgroup[i].Code, false))
			check(os.MkdirAll(subgroupOutDir, 0755))
			_, err = copy(pathToEmoji, path.Join(subgroupOutDir, subgroup[i].Description+".png"))
			check(err)

		}
	}
}
