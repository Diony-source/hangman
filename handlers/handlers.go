// Manages user interaction for the Hangman game
package handlers

import (
	"bufio"
	"fmt"
	"hangman/entities"
	"hangman/services"
	"os"
	"strings"
)

// StartGame initializes and runs the Hangman game
func StartGame() {
	game := entities.Game{
		Word:           "golang", // Secret word to guess
		Guessed:        []string{},
		RemainingTries: 6, // Allow 6 wrong guesses
	}

	reader := bufio.NewReader(os.Stdin)
	for !services.IsGameOver(game) {
		// Display the current state of the word
		fmt.Println("Current word:", services.DisplayWord(game))
		fmt.Printf("Remaining tries: %d\n", game.RemainingTries)

		// Get the user's guess
		fmt.Print("Enter a letter: ")
		letter, _ := reader.ReadString('\n')
		letter = strings.TrimSpace(letter)

		// Update the game state
		if services.UpdateGame(&game, letter) {
			fmt.Println("Good guess!")
		} else {
			fmt.Println("Wrong guess!")
		}
	}

	// Game over
	if strings.Contains(services.DisplayWord(game), "_") {
		fmt.Println("You lost! The word was:", game.Word)
	} else {
		fmt.Println("Congratulations! You guessed the word:", game.Word)
	}
}
