package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSum(dp [][]int, a, b, x, y int) int {
	return dp[x][y] + dp[a-1][b-1] - dp[x][b-1] - dp[a-1][y]
}

func check(side int, dp, dist [][]int, n, m, d int) bool {
	if side == 0 {
		return true
	}

	for i := 1; i <= n; i++ {
		nextI := i + side - 1
		if nextI > n {
			break
		}
		for j := 1; j <= m; j++ {
			nextJ := j + side - 1
			if nextJ > m {
				break
			}
			newD := d - 1

			a := max(1, i-newD)
			b := max(1, j-newD)
			x := min(n, nextI+newD)
			y := min(m, nextJ+newD)

			if getSum(dp, a, j, x, nextJ) != 0 || getSum(dp, i, b, nextI, y) != 0 {
				continue
			}

			mnD := min(dist[i][j], dist[nextI][nextJ])
			mnD = min(mnD, min(dist[i][nextJ], dist[nextI][j]))
			if mnD < d {
				continue
			}
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nmdStr, _ := reader.ReadString('\n')
	nmdParts := strings.Fields(nmdStr)
	n, _ := strconv.Atoi(nmdParts[0])
	m, _ := strconv.Atoi(nmdParts[1])
	d, _ := strconv.Atoi(nmdParts[2])

	v := make([][]byte, n+2)
	for i := range v {
		v[i] = make([]byte, m+2)
	}
	for i := 1; i <= n; i++ {
		line, _ := reader.ReadString('\n')
		for j := 1; j <= m; j++ {
			v[i][j] = line[j-1]
		}
	}

	INF := max(n, m)
	dist := make([][]int, n+2)
	dp := make([][]int, n+2)
	for i := range dist {
		dist[i] = make([]int, m+2)
		dp[i] = make([]int, m+2)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	for i := 1; i <= n; i++ {
		prev := 0
		for j := 1; j <= m; j++ {
			if v[i][j] == 'x' {
				prev++
			}
			dp[i][j] = prev + dp[i-1][j]
			if v[i][j] == 'x' {
				dist[i][j] = 0
			} else {
				mnD := min(dist[i-1][j], dist[i][j-1]) + 1
				dist[i][j] = mnD
			}
		}
	}

	for i := n; i > 0; i-- {
		for j := m; j > 0; j-- {
			mnD := min(dist[i+1][j], dist[i][j+1]) + 1
			dist[i][j] = min(dist[i][j], mnD)
		}
	}

	left, right := 0, min(n, m)
	for left < right {
		mid := left + (right-left+1)/2
		if check(mid, dp, dist, n, m, d) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Println(left)
}
