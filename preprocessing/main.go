package main

import (
	"github.com/smorken/freerange/preprocessing/emoji"
)

func main() {

	emoji.DownloadFile("./emoji/emoji-test.txt", "http://unicode.org/Public/emoji/11.0/emoji-test.txt")
	emoji.ParseEmojiData("./emoji/emoji-test.txt", "./emoji/emoji-test.json")
	emoji.ProcessEmojis("./emoji/emoji-test.json", "../frontend/assets")
	emoji.ProcessBackgroundEmojis("./emoji/backgrounds.json", "./emoji/emoji-test.json", "../frontend/assets")
}
