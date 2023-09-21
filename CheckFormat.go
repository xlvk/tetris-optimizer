package main

import "log"

// Checks file format
func CheckFormat(transformedContent [][]string) {
	var validhashcount int
	for _, tetromino := range transformedContent {
		for _, row := range tetromino {
			if len(row) != 4 {
				log.Fatal("ERROR")
			}
			for _, ch := range row {
				if ch == '#' {
					validhashcount++
				}
			}
		}
	}
	if validhashcount != 4 * len(transformedContent) {
		log.Fatal("ERROR")
	}
}
