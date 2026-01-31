package main

func SieveOfEratosthenes(n int) []int {
	mask := make([]bool, n-1)
	for i := range mask {
		if !mask[i] {
			for j := i + 1; j < n-1; j++ {
				if (j+2)%(i+2) == 0 {
					mask[j] = true
				}
			}
		}
	}
	var res []int
	for i, fl := range mask {
		if !fl {
			res = append(res, i+2)
		}
	}
	return res
}
