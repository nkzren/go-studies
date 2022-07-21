package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		acc := a
		b, a = a, acc+b
		return acc
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Printf("Element %d: %d\n", i+1, f())
	}
}
