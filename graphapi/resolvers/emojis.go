package resolvers

import (
	"context"

	"github.com/skimke/graphql-app/graphapi/data"
	"github.com/skimke/graphql-app/graphapi/gql"
)

type emojisresolver struct{}

// New will return a new todo resolver
func New() gql.Resolvers {
	return &emojisresolver{}
}

func (e *emojisresolver) Mutation_generateEmojis(ctx context.Context, times int, length int) ([]data.Emojis, error) {
	emojiStrings := make([]data.Emojis, times)
	for i := 0; i < times; i++ {
		newEmojiString := data.CreateStringOfEmojis(length)
		emojiStrings[i] = data.Emojis{Text: newEmojiString}
	}
	return emojiStrings, nil
}
