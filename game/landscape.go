package helicoptergame

import tl "github.com/JoelOtter/termloop"

type Landscape struct {
	*tl.Entity
	BackgroundRectange *tl.Rectangle
	Width              int
	Height             int
	LandscapeBorder    map[Coordinates]int
}

func NewLandscape(w, h int) *Landscape {
	landscape := new(Landscape)
	// Width and height of the arena are decresed by one to add corners on the arena border.
	landscape.Width = w - 1
	landscape.Height = h - 1
	// Each arena cell will have a width and heigth of 1.
	landscape.Entity = tl.NewEntity(1, 1, 1, 1)
	// Creates a map of coordinates.
	landscape.LandscapeBorder = make(map[Coordinates]int)
	landscape.BackgroundRectange = tl.NewRectangle(1, 1, landscape.Width-1, landscape.Height-1, tl.ColorBlue)

	// This for loop will create the top and bottom borders
	for x := 0; x < landscape.Width; x++ {
		landscape.LandscapeBorder[Coordinates{x, 0}] = 1
		landscape.LandscapeBorder[Coordinates{x, landscape.Height}] = 1
	}

	// This for loop will create the left and right borders
	for y := 0; y < landscape.Height+1; y++ {
		landscape.LandscapeBorder[Coordinates{0, y}] = 1
		landscape.LandscapeBorder[Coordinates{landscape.Width, y}] = 1
	}
	return landscape
}

// Contains checks if the arenaborder map contains the coordinates of the snake, if so this will return true.
func (landscape *Landscape) Contains(c Coordinates) bool {
	_, exists := landscape.LandscapeBorder[c]
	return exists
}

// Draw is a termloop function that will draw out the arena border, this is called when the game has started.
func (arena *Landscape) Draw(screen *tl.Screen) {
	// This for loop will range ArenaBorder containing the coordinates of the arenaborder and will print them out on the screen.
	for i := range arena.LandscapeBorder {
		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: tl.ColorWhite,
		})
	}
}
