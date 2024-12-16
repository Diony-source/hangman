package main

import (
	"fmt"
	"hangman/handlers"
)


func main() {
	fmt.Println("Welcome to the Hangman Game!")
	handlers.StartGame()
}