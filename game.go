package main

type TGame struct {
	GameModes []GameMode
}

type GameMode interface {
	Name() string
	Description() string
}
