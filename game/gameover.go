package helicoptergame

import (
	"fmt"
	"io/ioutil"
	"os"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

type Gameoverscreen struct {
	tl.Level
	Logo        *tl.Entity
	OptionsText []*tl.Text
}

func (ts *Gameoverscreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen(0)
			sg.Screen().SetLevel(gs)
		}
		if event.Key == tl.KeyDelete || event.Key == tl.KeyBackspace || event.Key == tl.KeyBackspace2 {
			tb.Close()
			os.Exit(0)
		}
	}
}

func Gameover(score int) {
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	logofile, _ := ioutil.ReadFile("util/gameover-logo.txt")
	gos.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))
	gos.OptionsText = []*tl.Text{
		tl.NewText(47, 13, fmt.Sprintf("You scored %d!", score), tl.ColorBlack, tl.ColorWhite),
		tl.NewText(47, 15, "Press \"Enter\" to restart!", tl.ColorBlack, tl.ColorWhite),
		tl.NewText(47, 17, "Press \"Delete\" to quit!", tl.ColorBlack, tl.ColorWhite),
	}

	gos.AddEntity(gos.Logo)

	for _, vv := range gos.OptionsText {
		gos.AddEntity(vv)
	}

	sg.Screen().SetLevel(gos)
}
