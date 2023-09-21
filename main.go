package main

import (
	"fmt"
	"os"
	// "math"
	// "io/ioutil"
	"strings"
	"log"
	"bufio"
)

func main() {
	// Read input from file
	if len(os.Args[:1]) != 1 {
		fmt.Println("ERROR.")
		return
	} else {
		// content, err := ioutil.ReadFile("exp/" + os.Args[1])
		// if err != nil {
		// 	fmt.Println("Error reading file:", err)
		// 	os.Exit(1)
		// }
		readFile, err := os.Open("exp/" + os.Args[1])
		defer readFile.Close()
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		content := ""
		for fileScanner.Scan() {
		// ha, e := strconv.Atoi(fileScanner.Text())
		if fileScanner.Text() == "" {
			content+="\n"
		} else {
			content+= fileScanner.Text() 
			content+="\n"
		}

	}

		// Process input content
		lines := strings.Split(string(content), "\n")
		// fmt.Println(lines)
		var transformedContent [][]string
		for u, line := range lines {
			if u != len(line)-1 {
				transformedContent = append(transformedContent, strings.Split(line, ""))
			}

		}

		// // Check the format of the input
		// CheckCheck(transformedContent)

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
