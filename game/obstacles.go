package helicoptergame

import (
	"math/rand"
	"time"

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

	NewX := RandomInsideArena(insideborderW, 1)
	NewY := RandomInsideArena(insideborderH, 1)
	obstacle.Rectangle = tl.NewRectangle(NewX, NewY, 1, 1, tl.ColorRed)
	obstacle.Position.X = NewX
	obstacle.Position.Y = NewY
	obstacle.SetPosition(obstacle.Position.X, obstacle.Position.Y)

	return obstacle
}

func RandomInsideArena(iMax int, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
