package stack

type stack struct {
	elements []string
}

func (stack *stack) empty() (result bool) {
	return len(stack.elements) == 0
}

func (stack *stack) push(newElement string) (element string) {
	stack.elements = append(stack.elements, newElement)
	element = newElement
	return element
}

func (stack *stack) pop() (element string) {
	length := len(stack.elements)
	if length > 0 {
		element = stack.elements[length-1]
		stack.elements = stack.elements[:length-1]
	}
	return element
}

func (stack *stack) peek() (element string) {
	length := len(stack.elements)
	if length > 0 {
		element = stack.elements[length-1]
	}
	return element
}
