package main

import (
	"testing"
)

func assertEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func Test_convert(t *testing.T) {
	assertEquals(t, "1", convert(1))
	assertEquals(t, "Fizz", convert(3))
	assertEquals(t, "Buzz", convert(5))
	assertEquals(t, "FizzBuzz", convert(15))
}
