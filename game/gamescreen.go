package helicoptergame

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	X         int
	Y         int
	Direction direction
}

type GameScreen struct {
	tl.Level
	LandscapeEntity *Landscape
	PlayerEntity    *Player
	Obstacles       []*Obstacle
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
	if player.SuccessCollision() {
		gs = NewGamescreen()
		sg.Screen().SetLevel(gs)
	}
	if player.BorderCollision() {
		// Calls the GameOver function to take the player to the game over screen.
		Gameover()
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

func NewGamescreen() *GameScreen {
	// Creates the gamescreen level and create the entities
	gs = new(GameScreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gs.LandscapeEntity = NewLandscape(70, 25)
	player := NewPlayer()

	gs.Obstacles = append(gs.Obstacles, NewObstacle(), NewObstacle(), NewObstacle(), NewObstacle(), NewObstacle(), NewObstacle())

	gs.AddEntity(gs.LandscapeEntity.BackgroundRectange)
	gs.AddEntity(gs.LandscapeEntity)
	gs.AddEntity(gs.LandscapeEntity.LandscapeGround)
	gs.AddEntity(player)

	for _, obstacle := range gs.Obstacles {
		gs.AddEntity(obstacle)
	}

	sg.Screen().SetFps(7)

	return gs
}
