package helicoptergame

import tl "github.com/JoelOtter/termloop"

func StartGame() {
	sg = tl.NewGame()
	sg.Screen().SetFps(30)

	ts = NewTitleScreen()

	ts.AddEntity(ts.Logo)

	for _, v := range ts.OptionsText {
		ts.AddEntity(v)
	}

	sg.Screen().SetFps(10)
	sg.Screen().SetLevel(ts)
	sg.Start()
}
