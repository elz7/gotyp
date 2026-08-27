package game

import (
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
)

type TypingSpeedGame struct {
	CorrectCount int
	WrongCount   int
	Words        []string
	Pos          int
	Green        *color.Color
	Red          *color.Color
}

func (g *TypingSpeedGame) PlayerMove(v *gocui.View, value string) {
	w := g.Words[g.Pos]
	var color *color.Color

	old := w
	if w[0] == '\n' {
		w = w[1:]
	}

	if w == value {
		color = g.Green
		g.CorrectCount++
	} else {
		color = g.Red
		g.WrongCount++
	}
	w = old

	color.Fprintf(v, "%s ", w)

	g.Pos++
}

func (g *TypingSpeedGame) Score() (correct int, wrong int) {
	return g.CorrectCount, g.WrongCount
}

func (g *TypingSpeedGame) GenerateGameData(b *gocui.View) {
	words := getWords()
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
			remainingSpace -= len(w)
		}
	}

	s := builder.String()
	g.Words = strings.Split(s, " ")
	b.WriteString(s)
	b.SetWritePos(0, 0)
}

func NewTypingSpeedGame() Game {
	return &TypingSpeedGame{Green: color.New(color.FgGreen), Red: color.New(color.FgRed)}
}
