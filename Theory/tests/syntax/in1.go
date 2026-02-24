package main

import "fmt"

type person struct {
	age int
}

type (
	InMemoryRepo interface{} // type, value
	Repository   interface{}
)

func New1() Repository {
	var e *InMemoryRepo
	return e // type = *InMemoryRepo, val = nil
}

func New2() Repository {
	return nil // type = nil, val = nil
}

func main() {
	people := make([]person, 2)       // {ptr 0x0A, len = 2, cap = 2}
	p1 := &people[1]                  // ptr 0x10 person{}
	p1.age++                          // ptr 0x10 person{age : 1}
	people = append(people, person{}) // {ptr 0x5B, len = 3, cap = 3}
	p1.age++                          // ptr 0x10 person{age : 2}
	fmt.Println(people[1].age)        // 1
	fmt.Println(p1.age)               // 2
	fmt.Println(people)               // [people{}, people{age : 1}, people{}]

	fmt.Println(New1() == New2()) // false

	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface == nil {
		println("It's nil!")
	}
}
