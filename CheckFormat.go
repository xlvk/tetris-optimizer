package main

import (
	"fmt"
	"os"
	"strings"
)

// Checks file format
func CheckFormat(transformedContent [][]string) {
	index := 0
	for i := 0; i < len(transformedContent); i++ {
		if len(strings.Join(transformedContent[i], "")) > 1 {
			if len(strings.Join(transformedContent[i], "")) != 5 {
			fmt.Println("Error. Bad Formatting.")
				os.Exit(1)
		}
		}
		for k := 0; k < len(transformedContent[i]); k++ {
			if transformedContent[i][k] == "#" {
				index++
			} else if transformedContent[i][k] == "." {

			} else if IsValid(transformedContent[i][k]) {
				fmt.Println("Error. Bad Formatting.", transformedContent[i][k])
				os.Exit(1)
			}
			if len(transformedContent[i]) == 1 && index != 4 {
				fmt.Println("Error. Bad Formatting || figure doesn't consist of 4 cubes.")
				os.Exit(1)
			} else if len(transformedContent[i]) == 1 && index == 4 { 
					index = 0
			}
		}
	}

}
