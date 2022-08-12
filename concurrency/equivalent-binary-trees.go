package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		if x != y || ok1 != ok2 {
			return false
		}
		if !ok1 || !ok2 {
			break
		}
	}
	return true
}

func main() {
	t1, t2 := tree.New(1), tree.New(2)
	fmt.Println(Same(t1, t2))
}
