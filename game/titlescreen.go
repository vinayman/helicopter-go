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

func NewTitleScreen() *TitleScreen {
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

func (ts *TitleScreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen(0)
			sg.Screen().SetLevel(gs)
		}
	}
	if event.Key == tl.KeyDelete || event.Key == tl.KeyBackspace || event.Key == tl.KeyBackspace2 {
		tb.Close()
		os.Exit(0)
	}
}
