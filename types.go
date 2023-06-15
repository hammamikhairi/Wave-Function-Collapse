package main

import "github.com/hajimehoshi/ebiten/v2"

type (
	Tile struct {
		// 200x200 px
		Image     *ebiten.Image
		Format    []string
		Adjacency *SideAdjacency
		Index     int
		Collapsed bool
	}

	Tiles []*Tile
)

type SideAdjacency struct {
	Up, Right, Down, Left []int
}

type Game struct {
	Tiles Tiles
	Grid  Tiles
}
