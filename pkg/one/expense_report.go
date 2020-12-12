package one

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	pathToData = "input/one.txt"
)

// FindExpenseKey expense report successfully
func FindExpenseKey() (int, error) {
	// path set as const in one.go
	data, err := ioutil.ReadFile(pathToData)
	if err != nil {
		return -1, err
	}
	expenses, err := ConvertByteDataToIntegerArray(data)
	if err != nil {
		return -1, err
	}
	sumPair, err := Find2020Sum(expenses)
	if err != nil {
		return -1, err
	}
	expenseCorrection := sumPair[0] * sumPair[1]
	return expenseCorrection, nil
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
