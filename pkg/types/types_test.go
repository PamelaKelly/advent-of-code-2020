package types

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name             string
		existingElements []int
		expectedElements []int
		input            int
	}{
		{
			name:             "push element to empty stack",
			existingElements: []int{},
			expectedElements: []int{
				1234,
			},
			input: 1234,
		},
		{
			name: "push element to non-empty stack",
			existingElements: []int{
				1234,
			},
			expectedElements: []int{
				1234,
				5678,
			},
			input: 5678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			// Instantiate stack and set existing elements
			stack := IntStack{
				Elements: tt.existingElements,
			}
			stack.Push(tt.input)
			if !reflect.DeepEqual(stack.Elements, tt.expectedElements) {
				t.Errorf("got stack with elements [%v] but expected stack with elements [%v]", stack.Elements, tt.expectedElements)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name             string
		existingElements []int
		expectedElements []int
	}{
		{
			name:             "pop element from empty stack",
			existingElements: []int{},
			expectedElements: []int{},
		},
		{
			name: "pop element from non-empty stack - edge case, stack only contains one element",
			existingElements: []int{
				1234,
			},
			expectedElements: []int{},
		},
		{
			name: "pop element from non-empty stack",
			existingElements: []int{
				1234,
				5678,
			},
			expectedElements: []int{
				1234,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			// Instantiate stack and set existing elements
			stack := IntStack{
				Elements: tt.existingElements,
			}
			stack.Pop()
			if !reflect.DeepEqual(stack.Elements, tt.expectedElements) {
				t.Errorf("got stack with elements %v but expected stack with elements %v", stack.Elements, tt.expectedElements)
			}
		})
	}
}
