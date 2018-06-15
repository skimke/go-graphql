package resolvers

import (
	"net/http"
	"testing"
	"unicode/utf8"
)

func TestMutation_generateEmojis(t *testing.T) {
	req, err := http.NewRequest("POST", "/query", nil)
	if err != nil {
		t.Fatal(err)
	}

	newResolver := emojisresolver{}
	emojiStrings, err := newResolver.Mutation_generateEmojis(req.Context(), 1, 1)

	for _, Emoji := range emojiStrings {
		if utf8.RuneCountInString(Emoji.Text) != 1 {
			t.Errorf("Expected run count in string of %v", 1)
		}
	}

	sizeOfEmojiStrings := len(emojiStrings)
	if sizeOfEmojiStrings != 1 {
		t.Errorf("Expected size of emojiStrings 1, but it was %d instead.", sizeOfEmojiStrings)
	}
}
