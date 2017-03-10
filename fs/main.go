package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, error := ioutil.ReadFile("test.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(string(data))
	}
}
