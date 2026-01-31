package Struct

type DQueue struct {
	back  *nodeDQ
	front *nodeDQ
}

type nodeDQ struct {
	value interface{}
	next  *nodeDQ
	prev  *nodeDQ
}

func (q *DQueue) PushFront(value interface{}) {
	newNode := &nodeDQ{value: value, next: q.front}
	if q.front == nil {
		q.back = newNode
		q.front = newNode
	} else {
		q.front.prev = newNode
		q.front = newNode
	}
}

func (q *DQueue) PushBack(value interface{}) {
	newNode := &nodeDQ{value: value, prev: q.back}
	if q.back == nil {
		q.back = newNode
		q.front = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}

func (q *DQueue) PopFront() interface{} {
	if q.front == nil {
		return nil
	}
	value := q.front.value
	q.front = q.front.next
	if q.front != nil {
		q.front.prev = nil
	} else {
		q.back = nil
	}
	return value
}

func (q *DQueue) PopBack() interface{} {
	if q.back == nil {
		return nil
	}
	value := q.back.value
	q.back = q.back.prev
	if q.back != nil {
		q.back.next = nil
	} else {
		q.front = nil
	}
	return value
}

func (q *DQueue) LookFront() interface{} {
	if q.front == nil {
		return nil
	}
	return q.front.value
}

func (q *DQueue) LookBack() interface{} {
	if q.back == nil {
		return nil
	}
	return q.back.value
}
