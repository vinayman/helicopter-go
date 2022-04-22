package helicoptergame

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

var insideborderW = 70 - 1
var insideborderH = 25 - 4

type Obstacle struct {
	*tl.Rectangle
	Position Coordinates
}

func NewObstacle() *Obstacle {
	obstacle := new(Obstacle)

	PositionX := RandomInsideArena(insideborderW, 1)
	PositionY := RandomInsideArena(insideborderH, 1)
	obstacle.Rectangle = tl.NewRectangle(PositionX, PositionY, 1, 1, tl.ColorRed)
	obstacle.Position = Coordinates{X: PositionX, Y: PositionY}

	return obstacle
}

func RandomInsideArena(iMax int, iMin int) int {
	return rand.Intn(iMax-iMin) + iMin
}
