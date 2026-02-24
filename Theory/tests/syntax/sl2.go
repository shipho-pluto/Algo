package main

import "fmt"

type Cart map[string]int

func (cart Cart) add(s string, i int) {
	cart[s] += i
}

func isNil(err error) bool {
	return err == nil
}

type er struct {
	msg string
}

func (e er) Error() string {
	//TODO implement me
	panic(e.msg)
}

func (e er) String() string {
	return e.msg
}

func main() {
	n := make([]int, 1, 3)
	fmt.Println(n)
	app(n, 1)
	fmt.Println(n)
	sl(n, []int{1, 2})
	fmt.Println(n)
	ind(n, 0, 2)
	fmt.Println(n)

	n = []int{1, 2, 3}
	add(n[0:2])
	fmt.Println(n)
	adds(n[0:2])
	fmt.Println(n)

	var c = make(Cart)
	var d = &c
	c.add("1", 1)
	(*d).add("1", 2)
	(*d)["1"] += 5

	fmt.Println(c)
	fmt.Println(*d)

	var err1 error           // -> {type = nil, data = nil}, bc error = interface
	fmt.Println(isNil(err1)) // true

	var err2 *er             // -> {type = nil, data = nil}, bc er = struct
	fmt.Println(isNil(err2)) // -> {type = er, data = nil}, bc er = struct (false)
	var e *er
	fmt.Println(e == nil) // true

	err2 = &er{msg: "hi"}    // -> {type = er, data = "hi"}, bc er = struct
	fmt.Println(isNil(err2)) // false

	err2 = nil               // -> {type = er, data = nil}, bc delete only data
	fmt.Println(isNil(err2)) // false

	err1 = err2              // -> {type = er, data = nil}, потому что мы положили типизированный интерфейс
	fmt.Println(isNil(err1)) // false

	err1 = nil               //, тк err1 остался интерфейсом, мы не присваивали значение, то заменили на nil тип
	fmt.Println(isNil(err1)) // true

}

func app(s []int, v int) {
	s = append(s, v)
}

func sl(s1, s2 []int) {
	copy(s1, s2)
}

func ind(s []int, i, v int) {
	s[i] = v
}

func add(n []int) {
	n = append(n, 1)
}

func adds(n []int) {
	n = append(n, 2, 2)
}
