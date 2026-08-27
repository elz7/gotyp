package game

import (
	"math/rand"

	"github.com/awesome-gocui/gocui"
)

type Game interface {
	PlayerMove(v *gocui.View, input string)
	GenerateGameData(b *gocui.View)
	Score() (correct int, wrong int)
}

var CurrentGame Game

func getWords() map[int][]string {
	switch Settings.Language {
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
	return nWords[rand.Intn(len(nWords))] + " "
}
