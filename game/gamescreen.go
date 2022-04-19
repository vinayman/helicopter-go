package helicoptergame

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	prevX     int
	prevY     int
	Direction direction
}

type GameScreen struct {
	tl.Level
	LandscapeEntity *Landscape
	PlayerEntity    *Player
}

func (player *Player) Draw(screen *tl.Screen) {
	player.Entity.Draw(screen)
}

func NewPlayer() *Player {
	player := new(Player)
	// Create a new entity for a 1x1 pixel.
	player.Entity = tl.NewEntity(5, 5, 1, 1)
	// Sets a standard direction to right, do not change this to up or left as the snake.
	// will crash into a wall right after the game starts.
	player.Direction = down

	return player
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
		if player.BorderCollision() {
			// Calls the GameOver function to take the player to the game over screen.
			Gameover()
		}
	}
}

func (player *Player) BorderCollision() bool {
	return gs.LandscapeEntity.Contains(Coordinates{player.prevX, player.prevY})
}

func NewGamescreen() *GameScreen {
	// Creates the gamescreen level and create the entities
	gs = new(GameScreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gs.LandscapeEntity = NewLandscape(70, 25)
	player := NewPlayer()
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '🚁'})

	gs.AddEntity(gs.LandscapeEntity)
	gs.AddEntity(gs.LandscapeEntity.BackgroundRectange)
	gs.AddEntity(player)

	sg.Screen().SetFps(10)

	return gs
}
