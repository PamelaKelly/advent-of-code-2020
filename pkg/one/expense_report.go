package one

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PamelaKelly/advent-of-code-2020/pkg/types"
)

const (
	pathToData      = "input/one.txt"
	sumPartsPartTwo = 3
)

// FindExpenseKey expense report successfully
func FindExpenseKey() (int, int, error) {
	// path set as const in one.go
	data, err := ioutil.ReadFile(pathToData)
	if err != nil {
		return -1, -1, err
	}
	expenses, err := ConvertByteDataToIntegerArray(data)
	if err != nil {
		return -1, -1, err
	}
	basic, err := Find2020Sum(expenses)
	if err != nil {
		return -1, -1, err
	}
	basicKey := basic[0] * basic[1]

	threePartElements, err := FindSum2020Stack(expenses)
	if err != nil {
		return -1, -1, err
	}

	threePartKey := threePartElements[0]
	for i := 1; i <= (len(threePartElements) - 1); i++ {
		threePartKey = threePartKey * threePartElements[i]
	}

	return basicKey, threePartKey, nil
}

// ConvertByteDataToIntegerArray ...
func ConvertByteDataToIntegerArray(data []byte) ([]int, error) {
	stringDataArray := strings.Split(string(data), "\n")
	intDataArray := []int{}
	for _, item := range stringDataArray {
		i, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string %s to integer with error: %v", item, err.Error())
		}
		intDataArray = append(intDataArray, i)
	}
	return intDataArray, nil
}

// Find2020Sum returns two entries that sum to the value of 2020, if present
func Find2020Sum(data []int) ([]int, error) {
	for i, num1 := range data {
		for j, num2 := range data {
			// skip current top level num but don't discount that there
			// could be two nums of same value at different locations
			if i == j {
				continue
			}
			if num1+num2 == 2020 {
				return []int{num1, num2}, nil
			}
		}
	}
	return nil, fmt.Errorf("no two numbers in the data provided equal 2020 when summed")
}

// FindSum2020Stack find the sum given the number of contributing parts,
// e.g. 2 numbers makes the sum, or 3 numbers make up the sum
// OK but doesn't work for actual data - missing some edge case
func FindSum2020Stack(data []int) ([]int, error) {
	for i := range data {
		found, elements := findSumForSection(data, i, sumPartsPartTwo)
		if found {
			return elements, nil
		}
	}
	return nil, fmt.Errorf("no valid combination in the data provided equal 2020 when summed")
}

// findSumForSection given an array of integers and a starting sum check if
// any combination of the starting point with any other two elements sums to 2020
// TODO: inefficient and only works for a depth of 3 - alternative recusive solution?
func findSumForSection(data []int, initial int, sumParts int) (bool, []int) {
	stack := types.IntStack{
		Elements: []int{data[initial]},
	}
	sum := data[initial]
	parts := 1
	for i, start := range data {
		if i != initial {
			stack.Push(start)
			parts++
			sum += start
			for j, next := range data {
				if i != j {
					stack.Push(next)
					parts++
					sum += next
					// happy path - correct sum parts and running sum
					if parts == sumParts && sum == 2020 {
						return true, stack.Elements
					}
					// Reached number of sum parts but total isn't correct
					// or has exceeded target sum - Undo last step
					if parts == sumParts && sum != 2020 || sum > 2020 {
						stack.Pop()
						parts--
						sum -= next
					}
				}
			}
			stack.Pop()
			parts--
			sum -= start
		}
	}
	return false, nil
}
