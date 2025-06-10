package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)

	n *= x
	m *= y
	v := make([][]byte, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		v[i] = []byte(scanner.Text())
	}

	count := 0
	mini := (x*y + 1) / 2

	for fl := 0; fl < n; fl += x {
		for ap := 0; ap < m; ap += y {
			res := 0
			for i := fl; i < fl+x; i++ {
				for j := ap; j < ap+y; j++ {
					if v[i][j] == 'X' {
						res++
					}
				}
			}
			if res >= mini {
				count++
			}
		}
	}

	fmt.Println(count)
}

func main() {
	solve()
}
