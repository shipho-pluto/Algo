package _struct

import (
	"fmt"
	"strconv"
)

type TreeInt struct {
	Value int
	Left  *TreeInt
	Right *TreeInt
}

type TreeCh struct {
	Value rune
	Left  *TreeCh
	Right *TreeCh
}

func (t *TreeInt) Print() {
	s := ""
	fmt.Println(recur(t, s))
}

func recur(t *TreeInt, s string) string {
	if t != nil {
		s += strconv.Itoa(t.Value) + recur(t.Left, s) + recur(t.Right, s)
	}
	return s
}
