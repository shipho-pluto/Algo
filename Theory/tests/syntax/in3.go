package main

import (
	"fmt"
	"io"
	"os"
)

type I interface {
	Foo()
}
type S struct{}

func (s *S) Foo() {
	fmt.Println("foo")
}

func Build() I {
	var res *S // type = nil, val = nil
	return res // type = *S, val = nil
}

type X struct {
	val int
}

func (x X) S() {
	fmt.Println(x.val)
}

func main() {
	i := Build()
	if i != nil {
		i.Foo()
	} else {
		fmt.Println("nil")
	}

	_ = io.Reader(os.Stdin)
	_ = io.Writer(os.Stdout)
	_ = io.Closer(os.Stdout)

	x := X{123}
	defer func() {
		x.S()
	}()
	x.val = 456
}
