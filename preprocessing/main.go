package main

import (
	"github.com/smorken/freerange/preprocessing/emoji"
)

func main() {
	emoji.ParseEmojiData("./emoji/emoji-test.txt", "output.json")

}
