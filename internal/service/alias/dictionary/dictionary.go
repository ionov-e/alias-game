package dictionary

import "context"

type Dictionary interface {
	List(ctx context.Context) ([]string, error)
	Word(ctx context.Context, number uint16) (string, error)
}
