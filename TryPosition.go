package main

import "os"

// Position selection start
func TryPosition(piece int, tetrominoesList []Tetrominoes, size int) {
	for y := 0; y < len(BOARD); y++ {
		for x := 0; x < len(BOARD); x++ {
			if CheckPosition(y, x, piece, tetrominoesList) {
				if y == len(BOARD)-1 || piece == len(tetrominoesList)-1 {
					PrintBoard()
					os.Exit(0)
				} else {
					TryPosition(piece+1, tetrominoesList, size)
				}
				ClearPosition(y, x, piece, tetrominoesList)
			}
		}
	}

	if piece == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		TryPosition(0, tetrominoesList, increaseSize)
	}
}
