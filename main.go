package main

import (
	"fmt"

	avl "app/avl"
)

func main() {
	tree := avl.Tree{Root: nil}
	tree.Insert(3)
	tree.Insert(4)
	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(3))
	fmt.Println(tree.Search(2))

}