package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		fmt.Println(convert(n))
	}
}

func convert(n int) string {
	isMultipleOfThree := n%3 == 0
	isMultipleOfFive := n%5 == 0
	if isMultipleOfThree && isMultipleOfFive {
		return "FizzBuzz"
	}
	if isMultipleOfThree {
		return "Fizz"
	}
	if isMultipleOfFive {
		return "Buzz"
	}
	return fmt.Sprint(n)
}
