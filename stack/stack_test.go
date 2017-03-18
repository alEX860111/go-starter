package stack

import (
	"testing"
)

func assertTrue(t *testing.T, value bool) {
	if !value {
		t.Errorf("Expected %v to be true", value)
	}
}

func assertStringEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func assertIntEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func Test_stack(t *testing.T) {
	stack := stack{}
	assertIntEquals(t, 0, len(stack.elements))
	assertTrue(t, stack.empty())
	assertStringEquals(t, "", stack.peek())

	assertStringEquals(t, "", stack.pop())
	assertIntEquals(t, 0, len(stack.elements))

	assertStringEquals(t, "joe", stack.push("joe"))
	assertStringEquals(t, "joe", stack.peek())
	assertStringEquals(t, "joe", stack.elements[0])
	assertIntEquals(t, 1, len(stack.elements))

	assertStringEquals(t, "zoe", stack.push("zoe"))
	assertStringEquals(t, "zoe", stack.peek())
	assertStringEquals(t, "joe", stack.elements[0])
	assertStringEquals(t, "zoe", stack.elements[1])
	assertIntEquals(t, 2, len(stack.elements))

	assertStringEquals(t, "zoe", stack.pop())
	assertStringEquals(t, "joe", stack.peek())
	assertStringEquals(t, "joe", stack.elements[0])
	assertIntEquals(t, 1, len(stack.elements))

	assertStringEquals(t, "joe", stack.pop())
	assertStringEquals(t, "", stack.peek())
	assertTrue(t, stack.empty())
	assertIntEquals(t, 0, len(stack.elements))
}
