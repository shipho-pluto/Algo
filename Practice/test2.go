package main

import "fmt"

func traceResult(A [][]int, weight []int, k, s int, result *[]int) {
	if A[k][s] == 0 {
		return
	}
	if A[k-1][s] == A[k][s] {
		traceResult(A, weight, k-1, s, result)
	} else {
		traceResult(A, weight, k-1, s-weight[k-1], result)
		*result = append(*result, k-1)
	}
}

func main() {
	weight := []int{3, 4, 5, 8, 9}
	prices := []int{1, 6, 4, 7, 6}
	count := len(prices)
	maxWeight := 13

	var A [][]int
	for i := 0; i < count+1; i++ {
		A = append(A, make([]int, maxWeight+1))
	}

	for k := 0; k < count+1; k++ {
		for s := 0; s < maxWeight+1; s++ {
			if k == 0 || s == 0 {
				A[k][s] = 0
			} else if s >= weight[k-1] {
				A[k][s] = max(A[k-1][s], A[k-1][s-weight[k-1]]+prices[k-1])
			} else {
				A[k][s] = A[k-1][s]
			}
		}
	}

	var result []int
	traceResult(A, weight, count, maxWeight, &result)
	fmt.Println(A[count][maxWeight], result)

}
