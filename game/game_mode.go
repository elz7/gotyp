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
		Name:        "Blind Typing",
		Description: "Several words will appear on the screen. Remember and retype them all",
	},
}
