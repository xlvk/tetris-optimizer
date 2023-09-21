package main

import (
	"fmt"
	"os"
)

// Checks file format
func CheckCheck(transformedContent [][]string) {
	index := ""
	for i := 0; i < len(transformedContent); i++ {
		// if transformedContent[i] != nil {
		for k := 0; k < len(transformedContent[i]); k++ {
			index += transformedContent[i][k]
			// fmt.Println(transformedContent[i][k], len(transformedContent[i]))
		}
		fmt.Println(index)
		if len(index) != 5 && index != "" {
			fmt.Println("Error. Bad Formatting.", index)
			os.Exit(1)
		} else {
			index = ""
		}
	}

	// }

}
