package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("t1 = ", t1)
	fmt.Println("t2 = ", t2)
	fmt.Println(Same(t1, t2))
}
