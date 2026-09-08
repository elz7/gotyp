package game

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
)

type TypingSpeedGame struct {
	score Score
	words []string
	pos   int
	green *color.Color
	red   *color.Color
}

func (g *TypingSpeedGame) PlayerMove(v *gocui.View, value string) bool {
	w := g.words[g.pos]
	var color *color.Color

	if w == "\n" {
		fmt.Fprint(v, "\n")
		g.pos++
		w = g.words[g.pos]
	}

	if w == value {
		color = g.green
		g.score.Correct++
	} else {
		color = g.red
		g.score.Wrong++
	}

	color.Fprintf(v, "%s ", w)
	g.pos++

	return g.pos == len(g.words)-1
}

func (g *TypingSpeedGame) Score() Score {
	return g.score
}

func (g *TypingSpeedGame) GenerateGameData(b *gocui.View) {
	g.words = make([]string, 0)
	g.pos = 0

	words := getWords(Settings.Language)
	x, y := b.Size()

	var builder strings.Builder
	for range y {
		remainingSpace := x
		for {
			w := getRandomWordThatFits(words, remainingSpace)
			builder.WriteString(w)
			g.words = append(g.words, w)
			if w == "\n" {
				break
			}
			builder.WriteString(" ")
			remainingSpace -= len(w) + 1
		}
	}

	b.SetWritePos(0, 0)
	b.WriteString(builder.String())
	b.SetWritePos(0, 0)

}

func NewTypingSpeedGame() Game {
	return &TypingSpeedGame{green: color.New(color.FgGreen), red: color.New(color.FgRed)}
}
