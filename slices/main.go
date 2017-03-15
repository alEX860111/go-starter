package main

import "fmt"

func main() {
	names := [5]string{"joe", "lisa", "zoe", "peter", "anna"}
	fmt.Println(names, len(names), cap(names))

	var first = names[0:3]
	printSlice(first, "first")

	var last = names[2:]
	printSlice(last, "last")

	names[2] = "ooops"

	printSlice(first, "first")
	printSlice(last, "last")

	letters := []string{"a", "b", "c"}
	printSlice(letters, "letters")
}

func printSlice(s []string, name string) {
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Printf("slice %s=%v has length=%d and capacity=%d \n", name, s, len(s), cap(s))
}
