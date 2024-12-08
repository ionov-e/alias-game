package dictionary

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	userDB "alias-game/internal/entity/user/db"
	"alias-game/internal/storage"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type RedisBaseRandom struct {
	key           dictionaryConstant.Key
	numberOfTries uint16
	db            storage.DictionaryDBInterface
}

func (b *RedisBaseRandom) word(ctx context.Context, wordNumber uint16, originalList func() []string) (string, error) {
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
func (b *RedisBaseRandom) listWordsStoreIfNeeded(ctx context.Context, originalList func() []string) ([]string, error) {
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

func (b *RedisBaseRandom) redisKey() userDB.DictionaryKeyAndTry {
	return userDB.NewDictionaryKeyAndTry(b.key, b.numberOfTries)
}

func (b *RedisBaseRandom) randomList(allWords []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // We shouldn't bother
	r.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })
	return allWords
}
