package main

import "fmt"

func Abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// count of many numbers sum % 3
	{
		var n int
		fmt.Scan(&n)

		k := make([]int, 3)
		for i := 0; i < n; i++ {
			var x int
			kCopy := make([]int, 3)
			fmt.Scan(&x)
			copy(kCopy, k)
			k[Abs1(x%3)]++
			for j := 0; j < 3; j++ {
				k[Abs1((x+j)%3)] += kCopy[j]
			}

		}
		fmt.Println(k[0])
	}

	var n int
	fmt.Scan(&n)
	k := make([]int, 3)
	for i := 0; i < n; i++ {
		var x int
		kCopy := make([]int, 3)
		fmt.Scan(&x)
		copy(kCopy, k)
		k[Abs1((x%3))]++
		if x%3 == 0 {
			k[0] += kCopy[0]
			k[1] += kCopy[1]
			k[2] += kCopy[2]
		} else if Abs1(x%3) == 1 {
			k[0] += kCopy[2]
			k[1] += kCopy[0]
			k[2] += kCopy[1]
		} else {
			k[0] += kCopy[1]
			k[1] += kCopy[2]
			k[2] += kCopy[0]
		}
	}

	fmt.Println(k[0])

	{ // sum of many numbers sum % 17
		var n int
		fmt.Scan(&n)
		k25 := make([]int, 17)
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			kCopy := make([]int, 17)
			copy(kCopy, k25)
			k25[Abs1(x%17)] = max(x, kCopy[Abs1(x%17)])
			for j := 0; j < 17; j++ {
				if k25[j] != 0 {
					k25[Abs1((x+j)%17)] = max(kCopy[j]+x, kCopy[Abs1((x+j)%17)])
				}
			}
		}

		fmt.Println(k25[0])
	}
}
