package dictionary

import (
	"alias-game/internal/database"
	dbConstants "alias-game/internal/database/constants"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type BaseRandomRedisDictionary struct {
	key           dbConstants.DictionaryKey
	numberOfTries uint16
	db            database.DB
}

func (b *BaseRandomRedisDictionary) DictionaryWord(ctx context.Context, wordNumber uint16, originalList func() []string) (string, error) {
	word, err := b.db.DictionaryWord(ctx, b.redisKey(), wordNumber)
	if err != nil {
		if wordNumber != 0 {
			return "", fmt.Errorf("error fetching word: %w", err)
		}

		words, err := b.listWordsStoreIfNeeded(ctx, originalList)
		if err != nil {
			return "", fmt.Errorf("error creating dictionary while fetching 1st word: %w", err)
		}
		return words[0], nil
	}
	return word, nil
}

// listWordsStoreIfNeeded returns a list of words for the given key, or creates a new one if it does not exist
func (b *BaseRandomRedisDictionary) listWordsStoreIfNeeded(ctx context.Context, originalList func() []string) ([]string, error) {
	words, err := b.db.DictionaryWordList(ctx, b.redisKey())
	if err == nil {
		return words, nil
	}

	newWords := b.randomList(originalList())
	err = b.db.DictionaryCreate(ctx, b.redisKey(), newWords)
	if err != nil {
		return nil, fmt.Errorf("error saving dictionary into storage: %w", err)
	}
	return newWords, nil
}

func (b *BaseRandomRedisDictionary) redisKey() dbConstants.DictionaryKeyAndTry {
	return dbConstants.NewDictionaryKey(b.key, b.numberOfTries)
}

func (b *BaseRandomRedisDictionary) randomList(allWords []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // We shouldn't bother
	r.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })
	return allWords
}

func RedisDictionaryFactoryMethod(keyAndTry dbConstants.DictionaryKeyAndTry, db database.DB) (Dictionary, error) {
	if keyAndTry.BaseKey == dbConstants.Easy1 {
		return NewEasy1(keyAndTry.TryNumber, db), nil
	}

	return nil, fmt.Errorf("unknown dictionary key: %s", keyAndTry.BaseKey)
}
