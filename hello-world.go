package main

import "fmt"

func main() {
	var joe = person{Name: "joe", Age: 12}
	var lisa = person{Name: "lisa", Age: 16}

	var persons = make([]person, 2)
	persons[0] = joe
	persons[1] = lisa
	for _, p := range persons {
		p.greet()
	}
}

func (p person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

type person struct {
	Name string
	Age  int
}
