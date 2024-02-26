package game

import (
	"hangman/difficulty"
	"slices"
	"strings"
)

type GameState int

const (
	NotStarted GameState = iota
	InProgress
	Win
	Loss
)

type Game struct {
	difficulty  difficulty.Difficulty
	word        string
	Tries       uint
	MaxTries    uint
	symbolsUsed []rune
	state       GameState
}

func NewGame() Game {
	return Game{state: NotStarted}
}

func (g *Game) Init(word string, tries uint) {
	g.word = word
	g.Tries = 0
	g.MaxTries = tries
	g.state = InProgress
}

func (g *Game) checkWin() {
	for _, char := range g.word {
		if !slices.Contains(g.symbolsUsed, char) {
			return
		}
	}
	g.state = Win
}

func (g *Game) checkLoss() {
	if g.Tries >= g.MaxTries {
		g.state = Loss
	}
}

func (g *Game) updateState() {
	g.checkWin()
	g.checkLoss()
}

func (g Game) State() GameState {
	return g.state
}

func (g *Game) AddGuess(char rune) bool {
	g.symbolsUsed = append(g.symbolsUsed, char)
	defer g.updateState()
	if !strings.ContainsRune(g.word, char) {
		g.Tries = g.Tries + 1
		return false
	}
	return true
}

func (g Game) WordMask() string {
	sb := strings.Builder{}
	for _, char := range g.word {
		if strings.ContainsRune(string(g.symbolsUsed), char) {
			sb.WriteRune(char)
		} else {
			sb.WriteRune('_')
		}
		sb.WriteRune(' ')
	}
	return sb.String()
}
