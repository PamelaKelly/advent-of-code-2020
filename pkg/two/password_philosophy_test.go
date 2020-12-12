package two

import (
	"os"
	"reflect"
	"testing"
)

func TestCountValidPasswords(t *testing.T) {
	tests := []struct {
		name     string
		input    []Entry
		expected int
	}{
		{
			name: "happy path - one valid password",
			input: []Entry{
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baade",
				},
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baaaade",
				},
			},
			expected: 1,
		},
		{
			name: "happy path - multiple valid passwords",
			input: []Entry{
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baade",
				},
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baaaade",
				},
				{
					Min:      2,
					Max:      7,
					Letter:   "a",
					Password: "baaaade",
				},
				{
					Min:      1,
					Max:      3,
					Letter:   "d",
					Password: "baaaadde",
				},
			},
			expected: 3,
		},
		{
			name: "happy path - edge case min zero",
			input: []Entry{
				{
					Min:      0,
					Max:      3,
					Letter:   "a",
					Password: "bde",
				},
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baaaade",
				},
			},
			expected: 1,
		},
		{
			name: "happy path - edge case max zero",
			input: []Entry{
				{
					Min:      0,
					Max:      0,
					Letter:   "a",
					Password: "bde",
				},
				{
					Min:      1,
					Max:      3,
					Letter:   "a",
					Password: "baaaade",
				},
			},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountValidPasswords(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %d valid passwords but got %d", tt.expected, actual)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name          string
		mockFilePath  string
		mockFileData  string
		expected      []Entry
		expectedError error
	}{
		{
			name:         "happy path - reads file and returns correct list of entries",
			mockFilePath: "input.txt",
			mockFileData: "2-6 c: fcpwjqhcgtffzlbj\n6-9 x: xxxtwlxxx\n2-3 g: gjggg\n",
			expected: []Entry{
				{
					Min:      2,
					Max:      6,
					Letter:   "c",
					Password: "fcpwjqhcgtffzlbj",
				},
				{
					Min:      6,
					Max:      9,
					Letter:   "x",
					Password: "xxxtwlxxx",
				},
				{
					Min:      2,
					Max:      3,
					Letter:   "g",
					Password: "gjggg",
				},
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock file and write each item to the file
			// as a new line to replicate the format of the input data
			f, _ := os.Create(tt.mockFilePath)
			defer os.Remove(tt.mockFilePath)
			f.WriteString(tt.mockFileData)
			actual, actualError := ParseInput(tt.mockFilePath)
			for i, a := range actual {
				t.Logf("Validating entry: %d", i)
				e := tt.expected[i]
				if a.Min != e.Min {
					t.Errorf("expected entry to have min %d but got %d", e.Min, a.Min)
				}
				if a.Max != e.Max {
					t.Errorf("expected entry to have max %d but got %d", e.Max, a.Max)
				}
				if a.Letter != e.Letter {
					t.Errorf("expected entry to have letter %s but got %s", e.Letter, a.Letter)
				}
				if a.Password != e.Password {
					t.Errorf("expected entry to have password %s but got %s", e.Password, a.Password)
				}
			}
			if reflect.TypeOf(actualError) != reflect.TypeOf(tt.expectedError) {
				t.Errorf("expected error %s but got %s", reflect.TypeOf(tt.expectedError), reflect.TypeOf(actualError))
			}
		})
	}
}
