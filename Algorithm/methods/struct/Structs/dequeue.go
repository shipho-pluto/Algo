package Structs

type DeQueue struct {
	data []int
	max  bool
}

func NewDeQueue(max bool) *DeQueue {
	return &DeQueue{
		data: []int{},
		max:  max,
	}
}

func (d *DeQueue) PutBack(val int) {
	d.data = append(d.data, val)
}

func (d *DeQueue) PopBack() {
	d.data = d.data[:len(d.data)-1]
}

func (d *DeQueue) TopBack() int {
	if len(d.data) == 0 {
		if d.max {
			return 1_000_000
		}
		return -1_000_000
	}
	return d.data[len(d.data)-1]
}

func (d *DeQueue) PutFront(val int) {
	newData := make([]int, len(d.data)+1)
	newData[0] = val
	copy(newData[1:], d.data)
	d.data = newData
}

func (d *DeQueue) PopFront() {
	newData := make([]int, len(d.data)-1)
	copy(newData, d.data[1:])
	d.data = newData
}

func (d *DeQueue) TopFront() int {
	if len(d.data) == 0 {
		if d.max {
			return 1_000_000
		}
		return -1_000_000
	}
	return d.data[0]
}

func (d *DeQueue) IsEmpty() bool {
	return len(d.data) == 0
}
