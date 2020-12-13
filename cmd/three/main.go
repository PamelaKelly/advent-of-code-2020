package main

import (
	"fmt"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/three"
)

func main() {
	// Day three
	trajectory, err := three.Run()
	if err != nil {
		fmt.Printf("failed to sled down the hill with error %s", err.Error())
	}
	fmt.Printf("Calculated Toboggan Trajectory: %d\n", trajectory)
}
