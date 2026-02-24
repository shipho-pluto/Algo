package main

import "fmt"

func main() {
	s1, s2 := "totocc", "gfgfyu"
	fmt.Println(isom(s1, s2))

	a := "afuuaaad"
	k := 2
	t := "au"
	fmt.Println(maxRepeatK(a, k))
	fmt.Println(onePer(a, t))

	arr := []rune{'R', 'L', 'R', 'U', 'U', 'R', 'D', 'L', 'U', 'R', 'D', 'L', 'R', 'R', 'R', 'U'}
	for _, r := range optimize(arr) {
		fmt.Print(string(r) + " ")
	}
	fmt.Println()

	ar := []int{1230, 99, 23001, 123, 111, 300021, 101010, 90000009, 9}
	fmt.Println(digitPerm(ar))

	nums := []int{1, 1, 1, 2, 3, 5, 5, 5}
	k = 2
	fmt.Println(mostKRepeat(nums, k))

	dictionary := []string{
		"i", "am", "have", "you", "we", "are", "siblings", "program", "language", "go", "work",
	}
	text := []string{
		"i", "am", "yuo", "programm", "languje", "go", "works",
	}

	fmt.Println(cntRightWords(dictionary, text))
}

func cntRightWords(dictionary []string, text []string) []string {
	st := make(map[string]bool, len(dictionary))
	for i := range dictionary {
		st[dictionary[i]] = true
	}

	var res []string
	for i := range text {
		if st[text[i]] {
			res = append(res, text[i])
		} else {
			for j := range text[i] {
				mayCor := text[i][:j] + text[i][j+1:]
				if st[mayCor] {
					res = append(res, text[i])
					break
				}
			}
		}
	}

	return res

}

func mostKRepeat(arr []int, k int) []int {
	var res []int
	st := make(map[int]int)
	mr := 0
	for i := range arr {
		st[arr[i]]++
		mr = max(mr, st[arr[i]])
	}
	bucket := make([][]int, mr+1)
	for i, v := range st {
		bucket[v] = append(bucket[v], i)
	}

	for i := len(bucket) - 1; i > -1; i-- {
		for j := range bucket[i] {
			res = append(res, bucket[i][j])
			k--
			if k == 0 {
				return res
			}
		}
	}
	return res
}

func digitPerm(ar []int) [][]int {
	var res [][]int
	st := make(map[[9]int][]int)
	for i := range ar {
		mask := [9]int{}
		c := ar[i]
		for c != 0 {
			if c%10 != 0 {
				mask[c%10-1]++
			}
			c /= 10
		}
		st[mask] = append(st[mask], ar[i])
	}

	for _, v := range st {
		res = append(res, v)
	}

	return res
}

func optimize(arr []rune) []rune {
	points := make(map[[2]int]bool)
	var res []rune
	x, y := 0, 0
	points[[2]int{0, 0}] = true
	for i := range arr {
		switch arr[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
		res = append(res, arr[i])
		if points[[2]int{x, y}] {
			xc, yc := x, y
			for j := len(res) - 1; j > -1; j-- {
				switch res[j] {
				case 'R':
					xc--
				case 'L':
					xc++
				case 'U':
					yc--
				case 'D':
					yc++
				}
				res = res[:j]
				points[[2]int{xc, yc}] = false
				if !points[[2]int{x, y}] {
					break
				}
			}
		}
		points[[2]int{x, y}] = true
	}

	return res
}

func onePer(a string, t string) int {
	mask := [26]int{}
	wind := [26]int{}
	for i := range t {
		mask[t[i]-'a']++
		wind[a[i]-'a']++
	}
	eq := 0
	for i := range 26 {
		if mask[i] == wind[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	for i := len(t); i < len(a); i++ {
		if mask[a[i-len(t)]-'a'] == wind[a[i-len(t)]-'a'] {
			eq--
		}
		wind[a[i-len(t)]-'a']--
		if mask[a[i-len(t)]-'a'] == wind[a[i-len(t)]-'a'] {
			eq++
		}

		if mask[a[i]-'a'] == wind[a[i]-'a'] {
			eq--
		}
		wind[a[i]-'a']++
		if mask[a[i]-'a'] == wind[a[i]-'a'] {
			eq++
		}

		if eq == 26 {
			return i - len(t) + 1
		}
	}

	return -1
}

func maxRepeatK(a string, k int) int {
	if len(a) <= k {
		return len(a)
	}
	res := 0
	mask := [26]int{}
	l := 0
	for r := range a {
		if mask[a[r]-'a'] == k {
			for l < r {
				if mask[a[r]-'a'] < k {
					break
				}
				mask[a[l]-'a']--
				l++
			}
		}
		res = max(res, r-l+1)
		mask[a[r]-'a']++
	}

	return res
}

func isom(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	mask := [26]int{}
	for i := range s1 {
		if mask[s1[i]-'a'] != 0 {
			if mask[s1[i]-'a'] != int(s1[i]-s2[i]+100) {
				return false
			}
		}
		mask[s1[i]-'a'] = int(s1[i] - s2[i] + 100)
	}

	return true
}
