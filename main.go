package main

import (
	"fmt"
	"github.com/adramalech/datastructures/trees"
	"math/rand"
)

func main() {
	var bst trees.Tree = trees.NewNode(15)

	var n int = 10000000
	var k int = 800000

	var i int = 1
	var num int

	for i < k {
		num = rand.Intn(n)

		if !bst.Exists(num) {
			bst.Insert(num)
			i++
		}
	}

	var count int = bst.Count()
	fmt.Printf("num = %count", count)

	bst = bst.Delete(num)

 	count = bst.Count()

	var outputStr string = bst.ToString()
	fmt.Println(outputStr)


}
