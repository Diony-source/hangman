// Defines core data structures for the Hangman game
package entities

// Game holds the current state of the game
type Game struct {
	Word          string   // The word to guess
	Guessed       []string // Correctly guessed letters
	RemainingTries int      // Number of tries left
}
