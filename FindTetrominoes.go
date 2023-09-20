package main

import (
	"fmt"
	"os"
)

// Finds all tetrominoes from the file
func FindTetrominoes(transformedContent [][]string) []Tetrominoes {
	var tetrominoToAppend Tetrominoes
	var tetrominoesList []Tetrominoes
	tetrisCounter := 1
	fffff := false
	for a := 0; a < len(transformedContent); a++ {
		for b := 0; b < len(transformedContent[a]); b++ {
			var found bool = false
			for i := 0; i < len(FIGURES); i++ {
				var skip bool = false
				for k := 0; k < len(FIGURES[i]); k++ {
					for m := 0; m < len(FIGURES[i][k]); m++ {
						if a+k == tetrisCounter*5-1 || b+m >= 4 {
							skip = true
							break
						}
						if transformedContent[a+k][b+m] == FIGURES[i][k][m] {
							continue
						}
						skip = true
						break
					}
					if skip {
						break
					}
				}
				if !skip {
					found = true
					fffff = found
					tetrominoToAppend.form = FIGURES[i]
					tetrominoesList = append(tetrominoesList, tetrominoToAppend)
					break
				}
			}
			if found {
				if a+(tetrisCounter*5-a) >= len(transformedContent) {
					a = len(transformedContent) - 1
				} else {
					a = a + (tetrisCounter*5 - a) - 1
				}
				fffff = found
				tetrisCounter++
				break
			}
		}
	}
	if !fffff {
		fmt.Println("Error. Unknown tetromino.")
		os.Exit(2)
	}

	return tetrominoesList
}
