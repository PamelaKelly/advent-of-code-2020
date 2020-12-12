package main

import (
	"fmt"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/two"
)

func main() {
	// Day 2
	validPasswords, err := two.Run()
	if err != nil {
		fmt.Printf("error getting valid passwords: %s", err.Error())
	}
	fmt.Printf("The number of valid passwords in the Toboggan system is: %d\n", validPasswords)
}
