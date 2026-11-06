package cui

import (
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
)

type Alignment int

const (
	AlignLeft    Alignment = 0
	AlignRight   Alignment = 1
	AlignCenter  Alignment = 3
	AlignJustify Alignment = 4
)

type ViewBufferStringOptions struct {
	Alignment Alignment
}

func changeViewVisibility(g *gocui.Gui, b bool, views ...string) {
	for _, it := range views {
		v, _ := g.View(it)
		v.Visible = b
	}
}

func SetViewBufferString(g *gocui.Gui, view, text string, options ...ViewBufferStringOptions) error {
	v, err := g.View(view)
	if err != nil {
		return err
	}
	v.Clear()

	if len(options) == 0 {
		fmt.Fprint(v, text)
		return nil
	}

	opts := options[len(options)-1]
	value := strings.Builder{}

	// Set alignment
	align := opts.Alignment
	switch align {
	case AlignLeft:
		break
	case AlignRight:
		width, _ := v.Size()
		for line := range strings.SplitSeq(text, "\n") {
			if l := len(line); l < width {
				value.WriteString(strings.Repeat(" ", width-l))
				value.WriteString(line)
			} else {
				value.WriteString(line)
			}
		}
	case AlignCenter:
		width, _ := v.Size()
		for line := range strings.SplitSeq(text, "\n") {
			if l := len(line); l < width {
				pad := (width - l) / 2
				value.WriteString(strings.Repeat(" ", pad))
				value.WriteString(line)
				value.WriteString(strings.Repeat(" ", pad))
			} else {
				value.WriteString(line)
			}
		}
	case AlignJustify:
		// i do not know what it is
		// so i will make text stretching
		width, _ := v.Size()
		for line := range strings.SplitSeq(text, "\n") {
			gaps := strings.Count(line, " ")
			space := (width - (len(line) - gaps)) / gaps
			value.WriteString(strings.ReplaceAll(line, " ", strings.Repeat(" ", space)))
			value.WriteString("\n")
		}
	}

	fmt.Fprint(v, value.String())
	return nil
}
