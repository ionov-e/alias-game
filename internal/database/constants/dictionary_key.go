package constants

type DictionaryKey string

const (
	Easy1 DictionaryKey = "e1"
)

func (d DictionaryKey) String() string {
	return string(d)
}
