package helicoptergame

import tl "github.com/JoelOtter/termloop"

var sg *tl.Game
var gs *GameScreen
var ts *TitleScreen

type direction int

const (
	up direction = iota
	down
)

type Coordinates struct {
	X int
	Y int
}
