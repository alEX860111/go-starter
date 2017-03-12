package main

import "fmt"

type counter struct {
	value int
}

func (c *counter) increment() {
	c.value = c.value + 1
}

func main() {
	c := counter{10}
	fmt.Printf("%+v\n", c)
	c.increment()
	fmt.Printf("%+v\n", c)
}
