package game

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
)

type TypingSpeedGame struct {
	score    Score
	words    []string
	index    int
	occupied int
}

func (g *TypingSpeedGame) PlayerMove(v *gocui.View, value string) bool {
	if g.index == len(g.words)-1 {
		return true
	}

	word := g.words[g.index]

	if word == "\n" {
		fmt.Fprint(v, "\n")
		g.index++
		word = g.words[g.index]
		g.occupied = 0
	}

	score := &g.score.Correct
	printer := color.New(color.FgGreen)
	if word != value {
		score = &g.score.Wrong
		printer = color.New(color.FgRed)
	}
	*score++

	width, _ := v.Size()
	if width-g.occupied == len(word)+1 {
		printer.Fprintf(v, " %s", word)
	} else {
		printer.Fprintf(v, "%s ", word)
	}
	g.occupied += len(word) + 1
	g.index++

	return false
}

func (g *TypingSpeedGame) Score() Score {
	return g.score
}

func (g *TypingSpeedGame) GenerateGameData(b *gocui.View) {
	g.words = make([]string, 0)
	g.index = 0
	g.occupied = 0

	words := getWords(Settings.Language)
	x, y := b.Size()

	var builder strings.Builder
	for range y {
		remainingSpace := x
		for remainingSpace > 0 {
			w := getRandomWordThatFits(words, remainingSpace)

			g.words = append(g.words, w)

			if remainingSpace-len(w) == 1 {
				builder.WriteString(" ")
				builder.WriteString(w)
				remainingSpace -= len(w) + 1
			} else {
				builder.WriteString(w)
				builder.WriteString(" ")
				remainingSpace -= len(w) + 1
			}
		}
		g.words = append(g.words, "\n")
		builder.WriteString("\n")
	}

	b.SetWritePos(0, 0)
	b.WriteString(builder.String())
	b.SetWritePos(0, 0)

}

func NewTypingSpeedGame() Game {
	return &TypingSpeedGame{}
}
