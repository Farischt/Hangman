package hangman

import "fmt"

func Greet() {
	fmt.Println(`
        		 _______ _       _______ _______ _______ _       
		|\     /(  ___  | (    /(  ____ (       |  ___  | (    /|
		| )   ( | (   ) |  \  ( | (    \/ () () | (   ) |  \  ( |
		| (___) | (___) |   \ | | |     | || || | (___) |   \ | |
		|  ___  |  ___  | (\ \) | | ____| |(_)| |  ___  | (\ \) |
		| (   ) | (   ) | | \   | | \_  ) |   | | (   ) | | \   |
		| )   ( | )   ( | )  \  | (___) | )   ( | )   ( | )  \  |
		|/     \|/     \|/    )_|_______)/     \|/     \|/    )_)`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(turns int) {
	var draw string
	switch turns {
	case 1:
		draw = `1`
	case 2:
		draw = `2`
	case 3:
		draw = `3`
	case 4:
		draw = `4`
	case 5:
		draw = `5`
	case 6:
		draw = `6`
	case 7:
		draw = `7`
	case 8:
		draw = `8`
	}
	fmt.Printf("\nTurns left %s \n", draw)
}

func drawState(g *Game, guess string) {
	fmt.Printf("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Printf("Used:")
	drawLetters(g.UsedLetters)

	switch g.State {
	case GoodGuess:
		fmt.Println(GoodGuess)
	case AlreadyGuess:
		fmt.Printf("%s Your guess: %s\n", AlreadyGuess, guess)
	case BadGuess:
		fmt.Print(BadGuess)
	case Win:
		fmt.Printf("%s The word was: %s\n", Win, guess)
		drawLetters(g.Letters)
	case Lose:
		fmt.Printf("%s The word was: %s\n", Lose, guess)
		drawLetters(g.Letters)

	}
}

func drawLetters(l []string) {
	for _, element := range l {
		fmt.Printf("%v", element)
	}
	fmt.Println()
}
