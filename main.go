package main

import (
	"fmt"
	"math/rand"

	"github.com/adramalech/datastructures/trees"
)

func main() {
	var bst trees.Tree = trees.NewNode(15)

	var n int = 10000

	var num int

	for i := 0; i < n; i++ {
		num = rand.Intn(n)

		bst.Insert(num)
	}

	var findNum int = bst.Find(10)

	fmt.Printf("found num = %d", findNum)
}
