package game

type GameMode struct {
	Name        string
	Description string
	CreateGame  func() Game
}

var GameModes = []GameMode{
	{
		Name:        "Typing speed Test",
		Description: "Try to type as fast as you can",
		CreateGame:  NewTypingSpeedGame,
	},
	{
		Name:        "Lyrics Game",
		Description: "Some cool description\nTo test justify alignment\nYo yo",
		CreateGame:  NewLyricsGame,
	},
}
