package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"

	hasLetter := LetterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. Got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"

	hasLetter := LetterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s. Got=%v", word, guess, hasLetter)
	}
}
