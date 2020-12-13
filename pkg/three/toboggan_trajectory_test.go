package three

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name          string
		mockFilePath  string
		mockFileData  string
		expected      map[int][]int
		expectedError error
	}{
		{
			name:         "happy path - forest with single row",
			mockFilePath: "input.txt",
			mockFileData: "...#..##.\n#..###...",
			expected: map[int][]int{
				0: []int{0, 0, 0, 1, 0, 0, 1, 1, 0},
				1: []int{1, 0, 0, 1, 1, 1, 0, 0, 0},
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Create(tt.mockFilePath)
			defer os.Remove(tt.mockFilePath)
			f.Write([]byte(tt.mockFileData))
			actual, actualError := ParseInput(tt.mockFilePath)
			for i, a := range actual {
				e := tt.expected[i]
				if !reflect.DeepEqual(a, e) {
					t.Errorf("Expected row %d to equal %v but got %v", i, e, a)
				}
			}
			if reflect.TypeOf(actualError) != reflect.TypeOf(tt.expectedError) {
				t.Errorf("Expected error %s but got %s", tt.expectedError.Error(), actualError.Error())
			}
		})
	}
}

func TestCountTreesOnSlope(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int][]int
		slope    Slope
		expected int
	}{
		{
			name: "happy path - correct number of trees counted",
			input: map[int][]int{
				0: []int{0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1},
				1: []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1},
				2: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1},
				3: []int{1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0},
			},
			slope: Slope{
				Across: 3,
				Down:   1,
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountTreesOnSlope(tt.input, tt.slope.Across, tt.slope.Down)
			if actual != tt.expected {
				t.Errorf("expected %d number of trees but got %d", tt.expected, actual)
			}
		})
	}
}
