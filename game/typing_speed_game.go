package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
)

type TypingSpeedGame struct {
	score struct {
		correct int
		wrong   int
	}
	words    []string
	index    int
	occupied int

	settings struct {
		// supported: en
		lang     string
		duration time.Duration
	}
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

	score := &g.score.correct
	printer := color.New(color.FgGreen)
	if word != value {
		score = &g.score.wrong
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

func (g *TypingSpeedGame) ScoreString() string {
	var (
		correct  = g.score.correct
		wrong    = g.score.wrong
		accuracy = float64(g.score.correct) / float64(g.score.correct+g.score.wrong) * 100.0
		rawWPM   = float64(g.score.correct+g.score.wrong) / g.settings.duration.Minutes()
		netWPM   = float64(g.score.correct) / g.settings.duration.Minutes()
	)
	return fmt.Sprintf("Correct: %d\nWrong: %d\nAccuracy: %02f\nRawWPM: %d\nNetWPM: %d",
		correct, wrong, accuracy, int(rawWPM), int(netWPM))
}

func (g *TypingSpeedGame) GenerateGameData(b *gocui.View) {
	g.words = make([]string, 0)
	g.index = 0
	g.occupied = 0

	words := getWords(g.settings.lang)
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

func (g *TypingSpeedGame) RunTimer(gui *gocui.Gui, view *gocui.View, syncDone chan struct{}) {
	dur := g.settings.duration

	for range time.Tick(time.Second) {
		title := formatDuration(dur)
		gui.Update(func(g *gocui.Gui) error {
			view.Title = title
			return nil
		})
		dur -= time.Second
		if dur.Seconds() < 0 {
			syncDone <- struct{}{}
			break
		}
	}
}

func NewTypingSpeedGame(language string, duration time.Duration) Game {
	game := TypingSpeedGame{}
	game.settings.lang = language
	game.settings.duration = duration
	return &game
}
