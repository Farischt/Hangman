package hangman

import (
	"strings"
	"testing"
)

func TestNewGameWithInvalidTurns(t *testing.T) {
	turns := 0
	word := "hello"
	_, err := New(turns, word)

	if err == nil {
		t.Errorf("Test should fail due to invalid turns when creating game")
	}
}

func TestNewGameWithTooShortWordToGuess(t *testing.T) {
	turns := 10
	word := "hf"
	_, err := New(turns, word)

	if err == nil {
		t.Errorf("Test should fail due to too short word to guess when creating game")
	}
}

// Test Guesses
func validState(t *testing.T, expectedState State, actualState State) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%s', got '%s'", expectedState, actualState)
		return false
	}
	return true
}

func TestGoodGuess(t *testing.T) {
	game, _ := New(5, "BOB")
	guess := "o"
	game.MakeGuess(guess)
	validState(t, GoodGuess, game.State)
}

func TestBadGuess(t *testing.T) {
	game, _ := New(5, "BOB")
	guess := "a"
	game.MakeGuess(guess)
	validState(t, BadGuess, game.State)
}

func TestAlreadyUsed(t *testing.T) {
	game, _ := New(5, "BOB")
	guess := "a"
	// Add a guess to the used letters
	game.UsedLetters = append(game.UsedLetters, strings.ToUpper(guess))
	game.MakeGuess(guess)
	validState(t, AlreadyGuess, game.State)
}

func TestGameWon(t *testing.T) {
	game, _ := New(5, "BOB")
	guess := "O"
	game.FoundLetters = []string{"B", " ", "B"}
	game.UsedLetters = append(game.UsedLetters, "B")
	game.MakeGuess(guess)
	validState(t, Win, game.State)
}

func TestGameLost(t *testing.T) {
	game, _ := New(5, "BOB")
	guess := "A"
	game.TurnsLeft = 1
	game.MakeGuess(guess)
	validState(t, Lose, game.State)
}
