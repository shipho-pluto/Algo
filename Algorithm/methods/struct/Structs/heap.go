package Structs

type Heap struct {
	data []int
}

func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	h.sortUp()
}

func (h *Heap) Pop() {
	h.data[0] = h.Bot()
	h.data = h.data[:len(h.data)-1]
	h.sortDown()
}

func (h *Heap) Top() int {
	return h.data[0]
}

func (h *Heap) Bot() int {
	return h.data[len(h.data)-1]
}

func (h *Heap) sortUp() {
	index := len(h.data) - 1
	for h.data[h.Parent(index)] < h.data[index] {
		h.data[h.Parent(index)], h.data[index] = h.data[index], h.data[h.Parent(index)]
		index = h.Parent(index)
	}
}

func (h *Heap) sortDown() {
	index := 0
	for {
		if h.data[index] < h.data[h.Left(index)] {
			h.data[h.Left(index)], h.data[index] = h.data[index], h.data[h.Left(index)]
			index = h.Left(index)
		} else if h.data[index] < h.Right(index) {
			h.data[h.Right(index)], h.data[index] = h.data[index], h.data[h.Right(index)]
			index = h.Right(index)
		} else {
			break
		}
	}
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Left(i int) int {
	return i*2 + 1
}

func (h *Heap) Right(i int) int {
	return i*2 + 2
}
