package three

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	pathToData = "input/three.txt"
)

// Run ...
func Run() (int, error) {
	forest, err := ParseInput(pathToData)
	if err != nil {
		return -1, err
	}
	return LetsGo(forest), nil
}

// LetsGo ...
func LetsGo(forest map[int][]int) int {
	trees := 0
	across, down := 0, 0
	// assuming all rows are the same length
	rowTrees := len(forest[0])
	heightForest := len(forest)
	for {
		across += 3
		down++
		// the edge of the map
		// when you reach the horizontal edge move across 3 but difference the length
		// to get back to the left hand side of the map - removes need to render full map
		// so 15 becomes 0
		// 16 becomes 1 etc
		if across >= rowTrees {
			across -= rowTrees
		}
		// the bottom of the map
		if down > heightForest-1 {
			fmt.Println("We've reached the bottom of the hill")
			return trees
		}
		if forest[down][across] == 1 {
			trees++
		}
	}
}

// ParseInput ...
// Returns a map representing a "forest"
// where the key indicates the row number in the matrix
// and the list of bools indicates whether a tree is present
// (true) or not (false)
func ParseInput(filepath string) (map[int][]int, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	forest := map[int][]int{}
	rows := strings.Split(string(data), "\n")
	for i, row := range rows {
		forest[i] = []int{}
		for _, char := range row {
			sq := fmt.Sprintf("%c", char)
			if sq == "." {
				forest[i] = append(forest[i], 0)
			} else if sq == "#" {
				forest[i] = append(forest[i], 1)
			}
		}
	}
	return forest, nil
}
