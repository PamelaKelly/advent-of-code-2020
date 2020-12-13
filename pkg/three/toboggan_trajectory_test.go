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
		expected      map[int][]bool
		expectedError error
	}{
		{
			name:         "happy path - forest with single row",
			mockFilePath: "input.txt",
			mockFileData: "...#..##.\n#..###...",
			expected: map[int][]bool{
				0: []bool{false, false, false, true, false, false, true, true, false},
				1: []bool{true, false, false, true, true, true, false, false, false},
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

func TestLetsGo(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int][]bool
		expected int
	}{
		{
			name: "happy path - correct number of trees counted",
			input: map[int][]bool{
				0: []bool{false, false, false, true, true, true, false, false, false, false, false, true, true, false, true},
				1: []bool{false, false, false, false, false, false, false, true, true, false, false, true, false, false, true},
				2: []bool{false, false, true, true, false, false, true, true, true, false, false, false, true, true, true},
				3: []bool{true, true, false, false, true, false, false, false, false, true, false, false, true, true, false},
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LetsGo(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %d number of trees but got %d", tt.expected, actual)
			}
		})
	}
}
