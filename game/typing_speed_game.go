package game

import (
	"log"
	"strings"
)

type TypingSpeedGame struct {
}

func (g *TypingSpeedGame) PlayerMove(value string) {
	log.Println(value)
}

func (g *TypingSpeedGame) GenerateGameData(x, y int) string {
	words := getWords()

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

	log.Println(builder.String())

	return builder.String()
}

func NewTypingSpeedGame() Game {
	return &TypingSpeedGame{}
}
