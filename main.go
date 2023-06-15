package main

import (
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	emptyImage    *ebiten.Image
	emptySubImage *ebiten.Image
)

func init() {
	emptyImage = ebiten.NewImage(3, 3)
	emptyImage.Fill(color.White)
	emptySubImage = ebiten.NewImageFromImage(emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image))
}

func (g *Game) Init() {
	g.Tiles = LoadTiles()

	// for _, tile := range g.Tiles {
	GridMap[g.Tiles[3].Index] = g.Tiles[3]
	GridMap[30] = &Tile{
		Image:     g.Tiles[4].Image,
		Format:    g.Tiles[4].Format,
		Adjacency: g.Tiles[4].Adjacency,
		Index:     30,
		Collapsed: true,
	}
	GridMap[45] = &Tile{
		Image:     g.Tiles[4].Image,
		Format:    g.Tiles[4].Format,
		Adjacency: g.Tiles[4].Adjacency,
		Index:     45,
		Collapsed: true,
	}

	PrepAdjencency(g.Tiles)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update() error {
	SpaceAvailable(g.Tiles)
	time.Sleep(time.Second * 0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tileWidth := SCREEN_WIDTH / GRID_WIDTH
	tileHeight := SCREEN_HEIGHT / GRID_WIDTH

	for _, tile := range GridMap {

		x := float64(tile.Index % GRID_WIDTH * tileWidth)
		y := float64(tile.Index / GRID_WIDTH * tileHeight)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileWidth)/(float64(TILE_SIZE)+float64(TILE_PADDING_X)), float64(tileHeight)/(float64(TILE_SIZE)+float64(TILE_PADDING_Y)))
		op.GeoM.Translate(x, y)

		screen.DrawImage(tile.Image, op)
	}

	if Gen < uint(GRID_SIZE) {
		SaveScreen(screen)
	}

}

func main() {
	game := &Game{}
	game.Init()

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("DIE !")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
