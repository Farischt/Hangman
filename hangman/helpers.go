package hangman

func LetterInWord(guess string, word []string) bool {
	for _, element := range word {
		if element == guess {
			return true
		}
	}
	return false
}

func hasUserWon(word []string, foundLetters []string) bool {
	for idx, element := range foundLetters {
		if element != word[idx] {
			return false
		}
	}
	return true
}
