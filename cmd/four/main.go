package main

import (
	"fmt"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/four"
)

func main() {
	res, err := four.Run()
	if err != nil {
		fmt.Printf("failed to process passports with error: %s", err.Error())
	}
	fmt.Printf("There are %d valid passports in the data provided\n", res)
}
