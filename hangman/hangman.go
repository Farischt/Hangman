package hangman

import (
	"errors"
	"strings"
)

type State string

const (
	GoodGuess    State = "Good guess !"
	AlreadyGuess State = "You already used this letter !"
	BadGuess     State = "Bad guess !"
	Win          State = "You win !"
	Lose         State = "You lose !"
)

type Game struct {
	State        State
	Letters      []string
	FoundLetters []string // Good guesses
	UsedLetters  []string
	TurnsLeft    int
}

// Instanciate a new game
// Takes the total number of turns and the word to guess
// Returns a pointer to the game and an error
func New(turns int, word string) (*Game, error) {

	if len(word) < 3 {
		return nil, errors.New("word must be at least 3 characters long")
	} else if turns <= 0 {
		return nil, errors.New("invalid number of turns")
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))

	for i := 0; i < len(letters); i++ {
		found[i] = " _ "
	}

	return &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}, nil
}

func (g *Game) MakeGuess(guess string) {
	guess = strings.ToUpper(guess)

	// Check if the letter is already used
	if LetterInWord(guess, g.UsedLetters) {
		g.State = AlreadyGuess
	} else if LetterInWord(guess, g.Letters) {
		g.State = GoodGuess
		// Reveal letter
		g.RevealLetter(guess)
		if hasUserWon(g.Letters, g.FoundLetters) {
			g.State = Win
		}
	} else {
		g.State = BadGuess
		g.BadLetter(guess)
		if g.TurnsLeft <= 0 {
			g.State = Lose
		}
	}
}

func (g *Game) RevealLetter(letterToReveal string) {
	g.UsedLetters = append(g.UsedLetters, letterToReveal)
	for idx, element := range g.Letters {
		if element == letterToReveal {
			g.FoundLetters[idx] = letterToReveal
		}
	}
}

func (g *Game) BadLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft--
}
