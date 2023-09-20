package main

// Clears certain position after apply
func ClearPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) {
	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if tetrominoesList[piece].form[a-y][b-x] == "." {
				continue
			}
			if BOARD[a][b] == string(65+piece) {
				BOARD[a][b] = "."
			}
		}
	}
}
