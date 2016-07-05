package main

import (
	"github.com/peterhellberg/emojilib"
)

func PrintEmoji(text string) string {
	return emojilib.ReplaceWithPadding(text)
}
