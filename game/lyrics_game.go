package game

import "github.com/awesome-gocui/gocui"

type LyricsGame struct {
}

// GenerateGameData implements [Game].
func (l *LyricsGame) GenerateGameData(b *gocui.View) {
	panic("unimplemented")
}

// PlayerMove implements [Game].
func (l *LyricsGame) PlayerMove(v *gocui.View, input string) bool {
	panic("unimplemented")
}

// Score implements [Game].
func (l *LyricsGame) ScoreString() string {
	panic("unimplemented")
}

func (l *LyricsGame) RunTimer(gui *gocui.Gui, view *gocui.View, syncDone chan struct{}) {
	panic("unimplemented")
}

func NewLyricsGame() Game {
	return &LyricsGame{}
}
