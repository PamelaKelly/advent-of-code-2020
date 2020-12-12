package one

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertByteDataToIntegerArray(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expected      []int
		expectedError error
	}{
		{
			name:  "happy path",
			input: []byte("1234\n5678"),
			expected: []int{
				1234,
				5678,
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			actual, actualErr := ConvertByteDataToIntegerArray(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected integer array %v but got integer array %v", tt.expected, actual)
			}
			if reflect.TypeOf(actualErr) != reflect.TypeOf(tt.expectedError) {
				t.Errorf("expected error type %v but got error type %v", reflect.TypeOf(actualErr), reflect.TypeOf(tt.expectedError))
			}
		})
	}
}

func TestFind2020Sum(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expected      []int
		expectedError error
	}{
		{
			name: "happy path",
			input: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			expected: []int{
				1721,
				299,
			},
			expectedError: nil,
		},
		{
			name: "happy path - edge case first num summed with self equals 2020 and should not be counted",
			input: []int{
				1010,
				1721,
				366,
				299,
				675,
				1456,
			},
			expected: []int{
				1721,
				299,
			},
			expectedError: nil,
		},
		{
			name: "happy path - edge case random num summed with self equals 2020 and should not be counted",
			input: []int{
				675,
				979,
				1010,
				299,
				1721,
				1456,
			},
			expected: []int{
				299,
				1721,
			},
			expectedError: nil,
		},
		{
			name: "failure path - no entry found",
			input: []int{
				675,
				979,
				1010,
				298,
				1721,
				1456,
			},
			expected:      nil,
			expectedError: fmt.Errorf("no two numbers in the data provided equal 2020 when summed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			actual, actualErr := Find2020Sum(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected integer array %v but got integer array %v", tt.expected, actual)
			}
			if reflect.TypeOf(actualErr) != reflect.TypeOf(tt.expectedError) {
				t.Errorf("expected error type %v but got error type %v", reflect.TypeOf(actualErr), reflect.TypeOf(tt.expectedError))
			}
		})
	}
}

func TestFind2020SumStack(t *testing.T) {
	tests := []struct {
		name          string
		inputData     []int
		expected      []int
		expectedError error
	}{
		{
			name: "happy path with sum parts of 3",
			inputData: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			expected: []int{
				979,
				366,
				675,
			},
			expectedError: nil,
		},
		{
			name: "happy path with sum parts of 3",
			inputData: []int{
				979,
				1721,
				366,
				299,
				675,
				1456,
			},
			expected: []int{
				979,
				366,
				675,
			},
			expectedError: nil,
		},
		{
			name: "happy path with sum parts of 3",
			inputData: []int{
				1721,
				366,
				299,
				675,
				1456,
				979,
			},
			expected: []int{
				366,
				675,
				979,
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := FindSum2020Stack(tt.inputData)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected integer array %v but got integer array %v", tt.expected, actual)
			}
			if reflect.TypeOf(actualErr) != reflect.TypeOf(tt.expectedError) {
				t.Errorf("expected error type %v but got error type %v", reflect.TypeOf(tt.expectedError), reflect.TypeOf(actualErr))
				t.Log(actualErr)
			}
		})
	}
}

func TestFind2020SumRecursive(t *testing.T) {
	type args struct {
		data     []int
		pos      int
		limit    int
		parts    int
		sum      int
		target   int
		elements []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "happy path - simple ordered example",
			args: args{
				data: []int{
					979,
					366,
					675,
				},
				pos:      0,
				limit:    3,
				parts:    0,
				sum:      0,
				target:   2020,
				elements: []int{},
			},
			expected: []int{
				979,
				366,
				675,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Find2020SumRecursive(tt.args.data, tt.args.pos, tt.args.limit, tt.args.parts, tt.args.sum, tt.args.target, tt.args.elements)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected elements %v but got elements %v", tt.expected, actual)
			}
		})
	}
}
