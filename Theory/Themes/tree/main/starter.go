package main

import (
	"Algoritm/Theory/Themes/tree"
	"fmt"
)

func main() {
	t := &tree.Node{Val: 50}
	t.AddElement(75, nil)
	t.AddElement(100, nil)
	t.AddElement(200, nil)
	t.AddElement(90, nil)
	t.AddElement(25, nil)
	t.FindRecursive()
	fmt.Println()
	t.RemoveElementOneSon(75, nil)
	t.FindRecursive()
	fmt.Println()
	t.RemoveElementNoSon(100, nil)
	t.FindRecursive()
}
