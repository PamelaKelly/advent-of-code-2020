package main

import (
	"fmt"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/one"
)

func main() {
	// First iteration solution
	result, err := one.FindExpenseKey()
	if err != nil {
		fmt.Printf("could not find expense key due to error: %s\n", err.Error())
	}
	fmt.Printf("The value needed to fix your expense report is: %d\n", result)
}
