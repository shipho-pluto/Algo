package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	sum := 0
	qStr, _ := reader.ReadString('\n')
	q := make([]int, n)
	qParts := strings.Fields(qStr)
	for i := 0; i < n; i++ {
		q[i], _ = strconv.Atoi(qParts[i])
		sum += q[i]
	}

	cStr, _ := reader.ReadString('\n')
	c := make([]int, n)
	cParts := strings.Fields(cStr)
	for i := 0; i < n; i++ {
		c[i], _ = strconv.Atoi(cParts[i])
	}

	abStr, _ := reader.ReadString('\n')
	abParts := strings.Fields(abStr)
	a, _ := strconv.Atoi(abParts[0])
	b, _ := strconv.Atoi(abParts[1])

	if a == b {
		fmt.Println(sum * a)
		return
	}
	ans := 0
	for i := 0; i < n; i++ {
		di := float64(c[i])*(float64(b-a)/255) + float64(a)
		ans += int(math.Ceil(float64(q[i]) * di))
	}
	fmt.Println(ans)
}
