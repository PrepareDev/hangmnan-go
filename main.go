package main

import (
	"fmt"
	"hangman/difficulty"
	"hangman/game"
	"log"
)

func main() {
	g := game.NewGame()
	diff, err := difficulty.SelectDifficulty()
	if err != nil {
		log.Fatal(err)
	}
	tries, err := difficulty.TriesFromDifficulty(diff)
	if err != nil {
		log.Fatal(err)
	}
	g.Init("hello", tries)

	if err != nil {
		log.Fatal(err)
	}

	for g.State() == game.InProgress {
		fmt.Println(g.WordMask())
		fmt.Println(g.Tries)
		var guess rune
		fmt.Scanf("%c\n", &guess)
		g.AddGuess(guess)
	}

	if g.State() == game.Loss {
		log.Fatal("YOU LOST")
	}
	if g.State() == game.Win {
		log.Fatal("YOU WON")
	}
}
