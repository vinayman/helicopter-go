package helicoptergame

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	X         int
	Y         int
	Direction direction
}

type GameScreen struct {
	tl.Level
	LandscapeEntity     *Landscape
	PlayerEntity        *Player
	Obstacles           []*Obstacle
	score               int
	FPS                 float64
	ObstacleCoordinates map[Coordinates]int
}

func (player *Player) Draw(screen *tl.Screen) {
	switch player.Direction {
	case down:
		player.Y++
		player.X++
	case up:
		player.Y--
		player.X++
		player.Direction = down
	}
	if player.BorderCollision() || player.ObstacleCollision() {
		Gameover(gs.score)
	}
	if player.SuccessCollision() {
		score := gs.score + 5
		gs = NewGamescreen(score)
		sg.Screen().SetLevel(gs)
	}

	screen.RenderCell(player.X, player.Y, &tl.Cell{
		Fg: tl.ColorGreen,
		Ch: '🚁',
	})
	player.Entity.Draw(screen)
}

func NewPlayer() *Player {
	player := new(Player)

	player.Direction = down
	player.X = 0
	player.Y = 12
	player.Entity = tl.NewEntity(player.X, player.Y, 1, 1)

	return player
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowUp:
			player.Direction = up
		case tl.KeyArrowDown:
			player.Direction = down
		}
	}
}

func (player *Player) BorderCollision() bool {
	return gs.LandscapeEntity.Contains(Coordinates{player.X, player.Y})
}

func (player *Player) SuccessCollision() bool {
	return gs.LandscapeEntity.SuccessContains(Coordinates{player.X, player.Y})
}

func (player *Player) ObstacleCollision() bool {
	_, exists := gs.ObstacleCoordinates[Coordinates{player.X, player.Y}]
	return exists
}

func NewGamescreen(score int) *GameScreen {
	gs = new(GameScreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gs.FPS = 6
	gs.score = score
	gs.LandscapeEntity = NewLandscape(70, 25)
	player := NewPlayer()

	gs.AddEntity(gs.LandscapeEntity.BackgroundRectange)
	gs.AddEntity(gs.LandscapeEntity)
	gs.AddEntity(gs.LandscapeEntity.LandscapeGround)
	gs.AddEntity(player)

	gs.ObstacleCoordinates = make(map[Coordinates]int)
	for i := 0; i < 6; i++ {
		obstacle := NewObstacle()
		gs.Obstacles = append(gs.Obstacles, obstacle)
		gs.AddEntity(obstacle)
		gs.ObstacleCoordinates[Coordinates{obstacle.Position.X, obstacle.Position.Y}] = 1
		gs.ObstacleCoordinates[Coordinates{obstacle.Position.X, obstacle.Position.Y - 1}] = 1
	}

	if score > 0 && score%5 == 0 {
		gs.FPS += 1
	}

	sg.Screen().SetFps(gs.FPS)

	return gs
}
