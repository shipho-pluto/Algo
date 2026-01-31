package main

import "fmt"

func main() {
	n := 38
	sum := 0
	for {
		for n != 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
		if sum < 10 {
			break
		}
		sum = 0

	}
	fmt.Println(sum)
}
