package main

import "fmt"

func main() {
	arr := "XXOOPSPOSSDFKDGPYFXGDFSYYRTYYYDFX"
	fmt.Println(nearXToY(arr))
}

func nearXToY(arr string) int {
	var isX bool
	n := len(arr)
	l := 0
	for ; l < n; l++ {
		if arr[l] == 'X' {
			isX = true
			break
		} else if arr[l] == 'Y' {
			isX = false
			break
		}
	}
	res := n
	for r := l + 1; r < n; r++ {
		if arr[r] == 'Y' && isX {
			res = min(res, r-l)
			l = r
			isX = false
		} else if arr[r] == 'X' && !isX {
			res = min(res, r-l)
			l = r
			isX = true
		} else if arr[r] == 'Y' && !isX {
			l = r
		} else if arr[r] == 'X' && isX {
			l = r
		}
	}

	return res
}
