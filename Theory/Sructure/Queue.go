package Struct

type Queue struct {
	back  *nodeQ
	front *nodeQ
}

type nodeQ struct {
	value interface{}
	next  *nodeQ
}

/*Пример
Допустим, очередь выглядит так:

[front] -> A -> B -> C <- [back]
Мы добавляем D:

newNode = D (пока никуда не ссылается).

q.back.next = D → теперь C ссылается на D:

[front] -> A -> B -> C -> D
                       ↑ back пока ещё указывает на C
q.back = D → теперь back указывает на D:

[front] -> A -> B -> C -> D <- [back]*/

// front сохраняет по итогу ссылку на следующий объект, а back постоянно меняется

func (q *Queue) Push(value interface{}) {
	newNode := &nodeQ{value: value}
	if q.back == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}

func (q *Queue) Pop() interface{} {
	if q.front == nil {
		return nil
	}
	value := q.front.value
	q.front = q.front.next
	return value
}
