package difficulty

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func TriesFromDifficulty(diff Difficulty) (uint, error) {
	switch diff {
	case Easy:
		return 10, nil
	case Medium:
		return 7, nil
	case Hard:
		return 5, nil
	}
	return 0, fmt.Errorf("unsupported difficulty passed: %v", diff)
}

func SelectDifficulty() (Difficulty, error) {
	var value Difficulty
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Difficulty]().Title("Select").Options(
				huh.NewOption("Easy", Easy),
				huh.NewOption("Medium", Medium),
				huh.NewOption("Hard", Hard),
			).Value(&value),
		),
	).Run()
	if err != nil {
		return 0, err
	}
	return value, nil
}
