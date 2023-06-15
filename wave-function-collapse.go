package main

import "math/rand"

func (ad *SideAdjacency) Add(side, index int) {
	switch side {
	case UP:
		ad.Up = append(ad.Up, index)

	case DOWN:
		ad.Down = append(ad.Down, index)

	case RIGHT:
		ad.Right = append(ad.Right, index)

	case LEFT:
		ad.Left = append(ad.Left, index)
	}
}

func (self *Tile) PrepRules(tiles Tiles) {

	for _, tile := range tiles {
		if Compare(self.Format[UP], tile.Format[DOWN]) {
			self.Adjacency.Add(UP, tile.Index)
		}

		if Compare(self.Format[RIGHT], tile.Format[LEFT]) {
			self.Adjacency.Add(RIGHT, tile.Index)
		}

		if Compare(self.Format[LEFT], tile.Format[RIGHT]) {
			self.Adjacency.Add(LEFT, tile.Index)
		}

		if Compare(self.Format[DOWN], tile.Format[UP]) {
			self.Adjacency.Add(DOWN, tile.Index)
		}
	}
	// fmt.Printf("%d >> %+v\n", self.Index, self.Adjacency)
}

func PrepAdjencency(tiles Tiles) {
	for _, tile := range tiles {
		tile.PrepRules(tiles)
	}
}

func SpaceAvailable(tiles Tiles) bool {

	for i := 0; i < GRID_SIZE; i++ {
		if _, ok := GridMap[i]; ok {
			continue
		}
		neighbours := GetNeighbours(i, tiles)
		possTiles := GetPossibleTiles(neighbours)
		if len(possTiles) == 0 {
			continue
		}

		parent := tiles[possTiles[rand.Intn(len(possTiles))]]
		GridMap[i] = &Tile{
			Image:     parent.Image,
			Format:    parent.Format,
			Adjacency: parent.Adjacency,
			Index:     i,
			Collapsed: true,
		}
		Gen++
		break
	}

	return false
}

func GetPossibleTiles(neighbours Tiles) []int {

	var (
		possUp, possRight, possDown, possLeft []int
		poss                                  = 0
		temp                                  int
	)

	if neighbours[0] != nil {
		possDown = neighbours[0].Adjacency.Down
		poss++
		temp = possDown[0]
	}

	if neighbours[1] != nil {
		possLeft = neighbours[1].Adjacency.Left
		poss++
		temp = possLeft[0]
	}
	if neighbours[2] != nil {
		possUp = neighbours[2].Adjacency.Up
		poss++
		temp = possUp[0]
	}
	if neighbours[3] != nil {
		possRight = neighbours[3].Adjacency.Right
		poss++
		temp = possRight[0]
	}

	possTiles := GetCommonTiles(possUp, possDown, possLeft, possRight)
	// fmt.Println(possTiles)
	// fmt.Printf("%+v\n", neighbours)
	// fmt.Println(possUp, possLeft, possRight, possDown)
	if poss == 1 {
		return []int{temp}
	}

	return possTiles
}

func GetNeighbours(index int, tiles Tiles) (neighbours Tiles) {

	neighbours = Tiles{nil, nil, nil, nil}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if i == j || (i == -1 && j == 1) || (i == 1 && j == -1) {
				continue
			}

			gridX := index%GRID_WIDTH + i
			gridY := index/GRID_WIDTH + j

			if gridX < 0 || gridX >= GRID_WIDTH || gridY < 0 || gridY >= GRID_WIDTH {
				continue
			}

			adjacentIndex := gridY*GRID_WIDTH + gridX
			neighbour, ok := GridMap[adjacentIndex]
			if !ok {
				continue
			}

			if i == -1 && j == 0 {
				neighbours[3] = neighbour // left
			} else if i == 0 && j == -1 {
				neighbours[0] = neighbour // up
			} else if i == 1 && j == 0 {
				neighbours[1] = neighbour
			} else {
				neighbours[2] = neighbour
			}
		}
	}

	return
}

func GetCommonTiles(slices ...[]int) []int {
	countMap := make(map[int]int)
	numSlices := 0
	for _, slice := range slices {
		for _, value := range slice {
			countMap[value]++
		}
	}

	var commonValues []int
	for _, slice := range slices {
		if len(slice) != 0 {
			numSlices++
		}
	}
	for value, count := range countMap {
		if count == numSlices {
			commonValues = append(commonValues, value)
		}
	}

	return commonValues
}
