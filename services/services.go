package services

import (
	"strings"
)

func UpdateGame(game *Game, letter string) bool {
	if strings.Contains(game.Word, letter) {
		if !strings.Contains(strings.Join(game.Guessed, ""), letter) {
			game.Guessed = append(game.Guessed, letter)
		}
		return true
	}

	game.RemainingTries--
	return false
}

func DisplayWord(game Game) string {
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

func IsGameOver(game Game) bool {
	return game.RemainingTries <= 0 || strings.Contains(DisplayWord(game), "_") == false
}
