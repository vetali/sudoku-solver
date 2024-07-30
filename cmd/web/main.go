package main

import (
	"fmt"
	"github.com/vetali/sudoku/pkg"
)

func main() {
	PrintRelatedCells()
}

func PrintRelatedCells() {
	for i := 0; i < 81; i++ {
		related := pkg.GetRelatedCells(i)
		fmt.Print("{")
		for idx, r := range related {
			if idx < len(related)-1 {
				fmt.Printf("%v, ", r)
			} else {
				fmt.Printf("%v", r)
			}
		}
		fmt.Printf("}, //%v\n", i)
	}
}
