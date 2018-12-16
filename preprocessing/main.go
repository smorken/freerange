package main

import (
	"github.com/smorken/freerange/preprocessing/emoji"
)

func main() {

	emoji.DownloadFile("./emoji/emoji-test.txt", "http://unicode.org/Public/emoji/11.0/emoji-test.txt")

	emoji.ParseEmojiData("./emoji/emoji-test.txt", "output.json")
	emoji.ProcessBackgroundEmojis("./emoji/backgrounds.json", "output.json", "../frontend/assets")
	emoji.ProcessEmojis("output.json", "../frontend/assets")
}
