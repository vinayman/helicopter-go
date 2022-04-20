package helicoptergame

import tl "github.com/JoelOtter/termloop"

type Landscape struct {
	*tl.Entity
	BackgroundRectange *tl.Rectangle
	Width              int
	Height             int
	LandscapeBorder    map[Coordinates]int
	DeadArea           map[Coordinates]int
	LandscapeGround    *tl.Rectangle
}

func NewLandscape(w, h int) *Landscape {
	landscape := new(Landscape)
	landscape.Width = w - 1
	landscape.Height = h - 1
	landscape.Entity = tl.NewEntity(1, 1, w, h)
	landscape.LandscapeBorder = make(map[Coordinates]int)
	landscape.DeadArea = make(map[Coordinates]int)
	landscape.BackgroundRectange = tl.NewRectangle(1, 1, landscape.Width-1, landscape.Height-1, tl.ColorBlue)
	landscape.LandscapeGround = tl.NewRectangle(1, 22, landscape.Width-1, 2, tl.ColorGreen)

	for x := 0; x < landscape.Width; x++ {
		landscape.LandscapeBorder[Coordinates{x, 0}] = 1
		landscape.LandscapeBorder[Coordinates{x, landscape.Height}] = 1
		landscape.DeadArea[Coordinates{x, landscape.Height - 3}] = 1
	}

	for y := 0; y < landscape.Height+1; y++ {
		landscape.LandscapeBorder[Coordinates{0, y}] = 1
		landscape.LandscapeBorder[Coordinates{landscape.Width, y}] = 1
	}

	return landscape
}

func (landscape *Landscape) Contains(c Coordinates) bool {
	_, exists := landscape.LandscapeBorder[c]
	if !exists {
		_, exists = landscape.DeadArea[c]
	}
	return exists
}

func (landscape *Landscape) Draw(screen *tl.Screen) {
	for i := range landscape.LandscapeBorder {
		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: tl.ColorWhite,
		})
	}
	for x := 1; x < landscape.Width-1; x++ {
		screen.RenderCell(x, landscape.Height-3, &tl.Cell{
			Fg: tl.ColorCyan, Ch: '🌳',
		})
	}
}
