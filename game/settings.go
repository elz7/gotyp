package game

import "time"

type GameSettings struct {
	Language      string
	RoundDuration time.Duration
}

var Settings = GameSettings{
	Language:      "en",
	RoundDuration: time.Minute,
}
