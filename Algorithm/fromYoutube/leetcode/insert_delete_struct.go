package main

import "math/rand"

type SomeStruct struct {
	data map[int]int
}

func (s *SomeStruct) Insert(value int) {
	s.data[value] = 1
}

func (s *SomeStruct) Delete(value int) {
	delete(s.data, value)
}

func (s *SomeStruct) GetRandom() int {
	key := rand.Intn(len(s.data))
	return s.data[key]
}
