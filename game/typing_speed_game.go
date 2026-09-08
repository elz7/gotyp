package game

import (
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

func (g *TypingSpeedGame) PlayerMove(v *gocui.View, value string) {
	w := g.words[g.pos]
	var color *color.Color

	old := w
	if w[0] == '\n' {
		w = w[1:]
	}

	if w == value {
		color = g.green
		g.score.Correct++
	} else {
		color = g.red
		g.score.Wrong++
	}
	w = old

	color.Fprintf(v, "%s ", w)

	g.pos++
}

func (g *TypingSpeedGame) Score() Score {
	return g.score
}

func (g *TypingSpeedGame) GenerateGameData(b *gocui.View) {
	words := getWords(Settings.Language)
	x, y := b.Size()

	var builder strings.Builder
	for range y {
		remainingSpace := x
		for {
			w := getRandomWordThatFits(words, remainingSpace)
			builder.WriteString(w)
			if w == "\n" {
				break
			}
			builder.WriteString(" ")
			g.words = append(g.words, w)
			remainingSpace -= len(w) + 1
		}
	}

	b.WriteString(builder.String())
	b.SetWritePos(0, 0)
}

func NewTypingSpeedGame() Game {
	return &TypingSpeedGame{green: color.New(color.FgGreen), red: color.New(color.FgRed)}
}
