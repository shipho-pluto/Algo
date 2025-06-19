package main

import "fmt"

func HowManyRooks(rooks [][]int) int {
	m := len(rooks)
	mI := make(map[int]int)
	mJ := make(map[int]int)
	for i := 0; i < m; i++ {
		mI[rooks[i][0]]++
		mJ[rooks[i][1]]++
	}
	cnt := 0
	for _, v := range mJ {
		cnt += v - 1
	}
	for _, v := range mI {
		cnt += v - 1
	}
	fmt.Println(mI, mJ)
	return cnt

}

func main() {
	rooks := [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {0, 7}, {3, 7}, {4, 5}, {7, 0}}
	fmt.Println(HowManyRooks(rooks))
}
