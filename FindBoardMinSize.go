package main

import "math"

// Calculates minimum possible board size
func FindBoardMinSize(tetrominoesList []Tetrominoes) int {
	minSideSize := 0
	for i := 0; i < len(tetrominoesList); i++ {
		for k := 0; k < len(tetrominoesList[i].form); k++ {
			if minSideSize < len(tetrominoesList[i].form[k]) {
				minSideSize = len(tetrominoesList[i].form[k])
			}

			if minSideSize < len(tetrominoesList[i].form) {
				minSideSize = len(tetrominoesList[i].form)
			}

			if minSideSize == 4 {
				break
			}
		}
	}

	blockCounter := math.Sqrt(float64(len(tetrominoesList) * 4))
	_, frac := math.Modf(blockCounter)
	if frac != 0 {
		blockCounter = math.Floor(blockCounter) + 1
	}

	return int(math.Max(blockCounter, float64(minSideSize)))
}
