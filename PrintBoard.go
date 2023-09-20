package main

import (
	"fmt"
	"strings"
)

// Print final board
func PrintBoard() {
	for _, row := range BOARD {
		fmt.Println(strings.Join(row, ""))
	}
}
