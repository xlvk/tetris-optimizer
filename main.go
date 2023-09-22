package main

import (
	"fmt"
	"os"

	// "math"
	"io/ioutil"
	"strings"
	// "tetrisOptimizer"
)

func main() {
	// Read input from file
	if len(os.Args[:1]) != 1 {
		fmt.Println("ERROR.")
		return
	} else {
		file := ""
		if !(strings.Contains(os.Args[1], "TestFile/")) {
			file = strings.TrimPrefix(os.Args[1], "TestFile/")
		} else {
			file = os.Args[1]
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		// Process input content
		lines := strings.Split(string(content), "\n")
		var transformedContent [][]string
		for _, line := range lines {
			transformedContent = append(transformedContent, strings.Split(line, ""))
		}

		// Check the format of the input
		CheckFormat(transformedContent)

		// Find the tetrominoes from the input
		tetrominoesList := FindTetrominoes(transformedContent)

		// Calculate the minimum board size
		size := FindBoardMinSize(tetrominoesList)

		// Create the board
		CreateBoard(size)

		// Try different positions to solve the puzzle
		TryPosition(0, tetrominoesList, size)

		fmt.Println("No solution found.")
	}

}

var BOARD [][]string

// All possible Tetris figures variations
var FIGURES [][][]string = [][][]string{
	{{"#", "#", "#"}, {".", "#", "."}},
	{{"#", "."}, {"#", "#"}, {"#", "."}},
	{{".", "#", "."}, {"#", "#", "#"}},
	{{"#", "."}, {"#", "#"}, {".", "#"}},
	{{"#", "#"}, {"#", "#"}},
	{{"#"}, {"#"}, {"#"}, {"#"}},
	{{"#", "#", "#", "#"}},
	{{"#", "#", "."}, {".", "#", "#"}},
	{{".", "#"}, {"#", "#"}, {"#", "."}},
	{{"#", "#", "."}, {".", "#", "#"}},
	{{".", "#"}, {"#", "#"}, {"#", "."}},
	{{"#", "."}, {"#", "."}, {"#", "#"}},
	{{"#", "#", "#"}, {".", ".", "#"}},
	{{"#", "#"}, {".", "#"}, {".", "#"}},
	{{"#", "."}, {"#", "#"}, {"#", "#"}},
	{{"#", "#"}, {"#", "#"}, {"#", "."}},
	{{".", ".", "#"}, {"#", "#", "#"}},
	{{"#", "#"}, {"#", "."}, {"#", "."}},
	{{"#", "#", "#"}, {"#", ".", "."}},
}

type Tetrominoes struct {
	form [][]string
}
