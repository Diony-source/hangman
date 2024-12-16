package entities

type Game struct {
	Word           string
	Guessed        []string
	RemainingTries int
}
