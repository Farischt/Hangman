package main

import (
	"fmt"
	dictionary "hangman/dictionnary"
	"hangman/hangman"
	"os"
)

func main() {
	err := dictionary.Load("./dictionnary/words.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	wordToGuess := dictionary.PickRandomWord()
	game, err := hangman.New(8, wordToGuess)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hangman.Greet()
	guess := ""
	for {
		hangman.Draw(game, guess)

		switch game.State {
		case hangman.Win:
			os.Exit(0)
		case hangman.Lose:
			os.Exit(1)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Couldn't read from the terminal: %e", err)
			os.Exit(1)
		}
		guess = l
		game.MakeGuess(guess)
	}
}
