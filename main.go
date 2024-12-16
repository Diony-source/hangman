// Entry point of the program
package main

import (
	"fmt"
	"hangman/handlers"
)

func main() {
	// Welcome message and start the game
	fmt.Println("Welcome to the Hangman Game!")
	handlers.StartGame()
}
