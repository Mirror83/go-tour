package main

// import (
// 	"golang.org/x/tour/tree"
// )

import "./tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// The walk is an
// inorder traversal of the tree
func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for valT1 := range ch1 {
		valT2 := <-ch2
		if valT1 != valT2 {
			return false
		}
	}

	return true
}
