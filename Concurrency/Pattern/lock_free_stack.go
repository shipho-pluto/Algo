package main

import (
	"sync/atomic"
	"unsafe"
)

type LockFreeStack struct {
	top unsafe.Pointer
}

type NodeLFS struct {
	val  int
	next unsafe.Pointer
}

func NewStack() LockFreeStack {
	return LockFreeStack{}
}

func (s *LockFreeStack) Push(val int) {
	next := &NodeLFS{val: val}

	for {
		head := atomic.LoadPointer(&s.top)
		next.next = head

		if atomic.CompareAndSwapPointer(&s.top, head, unsafe.Pointer(next)) {
			return
		}
	}

}

func (s *LockFreeStack) Pop() int {
	for {
		head := atomic.LoadPointer(&s.top)
		if head == nil {
			return -1
		}

		next := atomic.LoadPointer(&(*NodeLFS)(head).next)
		if atomic.CompareAndSwapPointer(&s.top, head, next) {
			return (*NodeLFS)(head).val
		}
	}
}
