package user

import "fmt"

type WordResult uint8

const (
	NotAnswered WordResult = iota
	Correct
	Incorrect
	Skipped
)

func (wr WordResult) String() string {
	switch wr {
	case NotAnswered:
		return "❓"
	case Correct:
		return "✅"
	case Incorrect:
		return "❌"
	case Skipped:
		return "❔"
	default:
		return fmt.Sprintf("%d", wr)
	}
}
