package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	// "sort"
	// "strconv"
)

func main() {
	var arr []string
	HashCount := 0
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR.")
		return
	} else {
		readFile, err := os.Open(args[0])
		defer readFile.Close()
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		line := fileScanner.Text()
		for fileScanner.Scan() {
			arr = append(arr, line)
			if len(arr[len(arr)-1])%4 == 0 && (len(arr)-1)%5 != 0 {
				for i := 0; i < len(line); i++ {
					if line[i] != '#' || line[i] != '.' {
						fmt.Println("ERROR.")
						return
					} else if line[i] == '#' {
						HashCount = HashCount + 1
					}
				}
			} else if (len(arr)-1)%5 == 0 {
				if line != "" {
					fmt.Println("ERROR.")
					return
				} else if HashCount != 4 {
					fmt.Println("ERROR.")
				return
				} else {
					arr[len(arr)-1] = "\n"
					HashCount = 0
				}
			} else {
				fmt.Println("ERROR.")
				return
			}
		}
	}
}
