package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	var valid = false
	for !valid {
		fmt.Printf("Enter your guess: ")
		guess, err = reader.ReadString('\n')

		if err != nil {
			return guess, err
		}

		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Println("Invalid letter input. ")
			continue
		}

		_, err := strconv.Atoi(guess)
		if err == nil {
			fmt.Println("Invalid letter input. Do not insert numbers !")
			continue
		}

		valid = true
	}

	return guess, nil
}
