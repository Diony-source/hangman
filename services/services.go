// Contains the logic for managing the Hangman game
package services

import (
	"hangman/entities"
	"strings"
)

// UpdateGame checks if the guessed letter is in the word and updates the game state
func UpdateGame(game *entities.Game, letter string) bool {
	if strings.Contains(game.Word, letter) {
		if !strings.Contains(strings.Join(game.Guessed, ""), letter) {
			game.Guessed = append(game.Guessed, letter)
		}
		return true
	}
	game.RemainingTries--
	return false
}

// DisplayWord shows the current state of the guessed word
func DisplayWord(game entities.Game) string {
	display := ""
	for _, char := range game.Word {
		letter := string(char)
		if strings.Contains(strings.Join(game.Guessed, ""), letter) {
			display += letter
		} else {
			display += "_"
		}
	}
	return display
}

// IsGameOver checks if the game is over
func IsGameOver(game entities.Game) bool {
	return game.RemainingTries <= 0 || !strings.Contains(DisplayWord(game), "_")
}
