package main

import (
	"sync/atomic"
	"unsafe"
)

type LockFreeQueue struct {
	tail unsafe.Pointer
	head unsafe.Pointer
}

type NodeLFQ struct {
	next unsafe.Pointer
	val  int
}

func MewQueue() LockFreeQueue {
	dummy := new(NodeLFQ)
	return LockFreeQueue{
		tail: unsafe.Pointer(dummy),
		head: unsafe.Pointer(dummy),
	}
}

func (q *LockFreeQueue) Push(val int) {
	node := &NodeLFQ{val: val}
	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*NodeLFQ)(tail).next)

		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*NodeLFQ)(tail).next, next, unsafe.Pointer(node)) {
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
					return
				}
			} else {
				atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
			}
		}
	}
}

func (q *LockFreeQueue) Pop() int {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*NodeLFQ)(head).next)

		if head == atomic.LoadPointer(&q.head) {
			if head == tail {
				if next == nil {
					return -1
				} else {
					atomic.CompareAndSwapPointer(&q.tail, tail, next)
				}
			} else {
				val := (*NodeLFQ)(next).val
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					return val
				}
			}
		}
	}
}
