package game

import "time"

type GameSettings struct {
	Language     string
	GameDuration time.Duration
}

var Settings = GameSettings{
	Language:     "en",
	GameDuration: 20 * time.Second,
}
