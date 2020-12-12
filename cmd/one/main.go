package main

import (
	"fmt"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/one"
)

func main() {
	// Day 1 - Part 1 - First iteration solution
	basicKey, threePartKey, err := one.FindExpenseKey()
	if err != nil {
		fmt.Printf("could not find expense key due to error: %s\n", err.Error())
	}
	fmt.Printf("The value needed to fix your expense report with 2 sum parts is: %d\n", basicKey)
	fmt.Printf("The value needed to fix your expense report with 3 sum parts is: %d\n", threePartKey)

}
