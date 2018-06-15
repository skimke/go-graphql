package data

import (
	"testing"
	"unicode/utf8"
)

func TestCreateStringOfEmojis(t *testing.T) {
	returnedString := CreateStringOfEmojis(3)
	runeCountInString := utf8.RuneCountInString(returnedString)

	if runeCountInString != 3 {
		t.Errorf("Expected run count of 3, but it was %d instead.", runeCountInString)
	}
}
