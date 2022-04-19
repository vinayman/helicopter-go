package helicoptergame

import (
	"io/ioutil"
	"os"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

type TitleScreen struct {
	tl.Level
	Logo        *tl.Entity
	OptionsText []*tl.Text
}

// NewTitleScreen will create a new titlescreen and return it.
func NewTitleScreen() *TitleScreen {
	// Create a title screen and its objects.
	ts = new(TitleScreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("util/titlescreen-logo.txt")
	ts.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))

	ts.OptionsText = []*tl.Text{
		tl.NewText(10, 15, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 17, "Press DELETE for options!", tl.ColorWhite, tl.ColorBlack),
	}

	return ts
}

// Tick will listen for a keypress to initiate the game.
func (ts *TitleScreen) Tick(event tl.Event) {
	// Checks if the event is a keypress event and the key pressed is the enter key.
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
	}
	if event.Key == tl.KeyDelete || event.Key == tl.KeyBackspace || event.Key == tl.KeyBackspace2 {
		tb.Close()
		os.Exit(0)
	}
}
