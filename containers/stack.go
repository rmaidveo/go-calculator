package containers

type Stack[T any] struct {
	elements []T
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Pop() (element T, ok bool) {
	if len(stack.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	lastElement := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]

	return lastElement, true
}
