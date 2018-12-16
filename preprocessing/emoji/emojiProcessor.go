package emoji

import (
	"encoding/json"
	"fmt"
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
		output := []map[string]string{}
		for i := 0; i < len(subgroup); i++ {
			pathToEmoji := path.Join("D:", "CODE", "twemoji", "2", "72x72", ProcessCode(subgroup[i].Code)+".png")
			if _, err := os.Stat(pathToEmoji); os.IsNotExist(err) {
				//file does not exist so dont use it
				continue
			}
			fmt.Printf("group[%s] name[%s] url[%s]\n", subgroupname, subgroup[i].Description, TwemojiURL(subgroup[i].Code, false))
			output = append(output,
				map[string]string{
					"name":    subgroup[i].Description,
					"png_url": TwemojiURL(subgroup[i].Code, false),
					"svg_url": TwemojiURL(subgroup[i].Code, true)})
		}
		jsonOutput, err := json.MarshalIndent(output, "", "    ")
		check(err)
		jsonOutputPath := path.Join(outdir, subgroupname+".json")
		err = ioutil.WriteFile(jsonOutputPath, jsonOutput, 0644)
		check(err)

	}
}
