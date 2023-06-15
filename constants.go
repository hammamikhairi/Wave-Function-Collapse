package main

const (
	TILES_DIR       string = "Tiles/"
	SCREENSHOTS_DIR string = "ScreenShots/"

	TILE_SIZE   int = 50
	TILE_BOUNDS int = 3

	TILE_PADDING_X int = 0
	TILE_PADDING_Y int = 0

	TILE_SIDES int = 4

	GRID_WIDTH int = 8
	GRID_SIZE  int = GRID_WIDTH * GRID_WIDTH

	SCREEN_HEIGHT int = GRID_WIDTH * 80
	SCREEN_WIDTH  int = GRID_WIDTH * 80
)

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

var ColorMap map[uint32]string = map[uint32]string{
	151: "A",
	0:   "B",
}

var TilesMap map[int]*Tile = map[int]*Tile{}
var GridMap map[int]*Tile = map[int]*Tile{}

var Gen uint = 0