package types

// Dice https://core.telegram.org/bots/api#dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based.
	Value int    `json:"value"` // Value of the dice, based on the emoji.
}
