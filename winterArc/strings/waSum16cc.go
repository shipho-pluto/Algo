package main

import "fmt"

func main() {
	n1, n2 := "fff", "1"
	fmt.Println(sum16cc(n1, n2))
}

func sum16cc(n1 string, n2 string) string {
	itoa := [16]int32{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	atoi := map[int32]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15}

	sum := ""
	add := '0'
	for i := 0; i < min(len(n2), len(n1)); i++ {
		el1 := int32(n1[len(n1)-1-i])
		el2 := int32(n2[len(n2)-1-i])
		sum += string(itoa[(atoi[add]+atoi[el1]+atoi[el2])%16])
		if (atoi[add] + atoi[el1] + atoi[el2]) >= 16 {
			add = itoa[(atoi[add]+atoi[el1]+atoi[el2])/16]
		} else {
			add = '0'
		}
	}
	for i := min(len(n2), len(n1)); i < max(len(n2), len(n1)); i++ {
		var el int32
		if len(n2) > len(n1) {
			el = int32(n2[len(n2)-1-i])
		} else {
			el = int32(n1[len(n1)-1-i])
		}
		sum += string(itoa[(atoi[add]+atoi[el])%16])
		if (atoi[add] + atoi[el]) >= 16 {
			add = itoa[(atoi[add]+atoi[el])/16]
		} else {
			add = '0'
		}
	}
	if add != '0' {
		sum += string(add)
	}
	ans := ""
	for i := len(sum) - 1; i > -1; i-- {
		ans += string(sum[i])
	}

	return ans
}
