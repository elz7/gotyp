package game

import (
	"math/rand"

	"github.com/awesome-gocui/gocui"
)

type Score struct {
	Correct int
	Wrong   int
}

type Game interface {
	// Returns true if there is a need for generating new game data, false otherwise.
	PlayerMove(v *gocui.View, input string) bool
	GenerateGameData(b *gocui.View)
	Score() Score
}

var CurrentGame Game

func getWords(lang string) map[int][]string {
	switch lang {
	case "en":
		return englishWords
	default:
		return englishWords
	}
}

func getMaxKey(v map[int][]string) int {
	max := 0
	for k := range v {
		if k > max {
			max = k
		}
	}
	return max
}

func getRandomWordThatFits(words map[int][]string, remainingSpace int) string {
	if remainingSpace <= 0 {
		return "\n"
	}

	i := min(rand.Intn(remainingSpace)+1, getMaxKey(words))
	nWords := words[i]
	return nWords[rand.Intn(len(nWords))]
}
