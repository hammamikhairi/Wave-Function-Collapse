package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"wfc/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadTiles() Tiles {
	tiles := Tiles{}
	store := assets.Store

	dir, err := store.ReadDir(TILES_DIR)
	if err != nil {
		log.Fatal(err)
	}

	for index, file := range dir {
		tileFile := filepath.Join(TILES_DIR, file.Name())

		buffer, err := store.Open(tileFile)
		if err != nil {
			log.Fatal(err)
		}

		tileImg, err := png.Decode(buffer)
		if err != nil {
			log.Fatal(err)
		}

		// TODO : ROTATE Tiles <REMOVE DUPS>

		format := TileFormat(tileImg)
		ebitenImg := ebiten.NewImageFromImage(tileImg)

		tiles = append(tiles, &Tile{
			Image:     ebitenImg,
			Format:    format,
			Index:     index,
			Adjacency: &SideAdjacency{},
		})
	}
	return tiles
}

func TileFormat(tileImg image.Image) []string {

	tileW := tileImg.Bounds().Max.X

	tran := tileW / (TILE_BOUNDS * 2)
	gap := tileW / TILE_BOUNDS

	var (
		side  string   = ""
		sides []string = []string{}
	)

	for i := 0; i < TILE_BOUNDS; i++ {
		side += GetColorKey(tileImg.At(gap*i+tran, 0))
	}
	sides = append(sides, side)

	side = ""
	for i := 0; i < TILE_BOUNDS; i++ {
		side += GetColorKey(tileImg.At(tileW-1, gap*i+tran))
	}
	sides = append(sides, side)

	side = ""
	for i := TILE_BOUNDS - 1; i >= 0; i-- {
		side += GetColorKey(tileImg.At(gap*i+tran, tileW-1))
	}
	sides = append(sides, side)

	side = ""
	for i := TILE_BOUNDS - 1; i >= 0; i-- {
		side += GetColorKey(tileImg.At(0, gap*i+tran))
	}
	sides = append(sides, side)

	return sides
}

func GetColorKey(rgba color.Color) string {
	r, _, _, _ := rgba.RGBA()
	return ColorMap[r>>8]
}

func ReverseString(side string) (res string) {
	for _, v := range side {
		res = string(v) + res
	}
	return
}

func Compare(a, b string) bool {
	return a == ReverseString(b)
}

func SaveScreen(screen *ebiten.Image) {
	img := ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT)
	img.DrawImage(screen, nil)
	buffer, err := os.Create(fmt.Sprintf("%sscreenshot%d.png", SCREENSHOTS_DIR, Gen))
	if err != nil {
		log.Fatal(err)
	}
	defer buffer.Close()

	err = png.Encode(buffer, img)
	if err != nil {
		log.Fatal(err)
	}
}
