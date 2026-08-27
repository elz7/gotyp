package cui

import (
	"io"
	"log"

	"github.com/awesome-gocui/gocui"
)

func initDebugConsole(w io.Writer) {
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.SetOutput(w)

	log.Println("[INF]: Debug Console is initialized.")
}

func debugPromptEnter(g *gocui.Gui, v *gocui.View) error {
	prompt := v.Buffer()
	defer v.Clear()

	switch prompt {
	case "/exit":
		return gocui.ErrQuit
	default:
		log.Println("[ERR] Unknown command:", prompt)
	}

	return nil
}
