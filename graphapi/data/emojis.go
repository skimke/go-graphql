package data

import (
	"math/rand"
)

type Emojis struct {
	Text string
}

var emojisList = []string{"\U0001f487", "\U0001f63b", "\U0001f36f", "\U0001f428", "\U0001f34b", "\U0001f913", "\U0001f30a", "\U0001f44c", "\U0001f989", "\U0001f6a2", "\U0001f916", "\U0001f64a", "\u2747\ufe0f", "\U0001f375", "\U0001f60b", "\U0001f680"}

func getRandom() string {
	n := rand.Intn(len(emojisList))

	return emojisList[n]
}

func CreateStringOfEmojis(length int) string {
	emojiString := ""
	for i := 0; i < length; i++ {
		emojiString = emojiString + getRandom()
	}

	return emojiString
}
