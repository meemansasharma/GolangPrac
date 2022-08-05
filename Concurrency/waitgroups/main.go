package main

import (
	"fmt"
	"sync"
)

func printArgument(value string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(value)
}
func main() {
	var wg sync.WaitGroup

	ar := []string{

		"apples",
		"grapes",
		"papaya",
		"pineapple",
		"apricot",
		"blue berry",
		"strawberry",
	}

	for i, x := range ar {

		wg.Add(1)
		go printArgument(fmt.Sprintf(" The value at index : %d is %s", i, x), &wg)

	}

	wg.Wait()
	wg.Add(1)
	printArgument("Last string", &wg)
}
