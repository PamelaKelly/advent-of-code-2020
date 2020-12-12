package types

// A place for shared types that might be useful to multiple packages

// IntStack is a struct that behaves like a typical Stack
// and is implemented using slices in Go
// Elements are integers
type IntStack struct {
	Elements []int
}

// Push adds the element provided to the top of the stack
func (r *IntStack) Push(element int) {
	r.Elements = append(r.Elements, element)
}

// Pop removes the item on the top of the stack
func (r *IntStack) Pop() {
	// Remove the last element by redefining the slice without the last element
	// This is slow but we need to maintain the order of the stack
	if len(r.Elements) == 1 {
		// only one element left so reset stack
		r.Elements = []int{}
	} else if len(r.Elements) > 1 {
		// more than one element return sub slice
		r.Elements = r.Elements[:len(r.Elements)-1]
	}
}
