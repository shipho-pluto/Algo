package main

import "fmt"

func main() {
	s := "OOOXOOOOXXXOOOOYYYYYOOOX"
	fmt.Println(minDistanceXY(s))

	points := [][]int{{2, 1}, {2, 0}, {0, 0}, {0, 1}, {2, 1}, {0, 1}}
	fmt.Println(lineReflection(points))

	s1, s2 := "totocc", "dadorr"
	fmt.Println(isomorph(s1, s2))

	arr := []int{4, 2, 2, 2, 4, 4, 2, 2}
	lim := 0
	fmt.Println(longestDiffMELim(arr, lim))
}

type DQ []int

func (s *DQ) PutFront(el int) {
	*s = append(*s, el)
}
func (s *DQ) PopFront() int {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}
func (s *DQ) LookFront() int {
	return (*s)[len(*s)-1]
}
func (s *DQ) PopBack() int {
	el := (*s)[0]
	*s = (*s)[1:]
	return el
}
func (s *DQ) LookBack() int {
	return (*s)[0]
}
func (s *DQ) Empty() bool {
	return len(*s) == 0
}
func longestDiffMELim(arr []int, lim int) any {
	dqMax, dqMin := DQ{}, DQ{}
	res := 0
	dqMin.PutFront(arr[0])
	dqMax.PutFront(arr[0])
	l, r := 0, 1
	for r < len(arr) {
		if dqMax.LookBack()-dqMin.LookBack() > lim {
			if arr[l] == dqMin.LookBack() {
				dqMin.PopBack()
			}
			if arr[l] == dqMax.LookBack() {
				dqMax.PopBack()
			}
			l++
		} else {
			for !dqMax.Empty() && dqMax.LookFront() < arr[r] {
				dqMax.PopFront()
			}
			dqMax.PutFront(arr[r])
			for !dqMin.Empty() && dqMin.LookFront() > arr[r] {
				dqMin.PopFront()
			}
			dqMin.PutFront(arr[r])
			res = max(res, r-l)
			r++
		}
	}
	if dqMax.LookBack()-dqMin.LookBack() <= lim {
		res = max(res, r-l)
	}
	return res
}

func isomorph(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	st := make(map[uint8]uint8)
	for i := range s1 {
		if v, ok := st[s1[i]]; ok {
			if v != s1[i]-s2[i]+26 {
				return false
			}
		}
		st[s1[i]] = s1[i] - s2[i] + 26
	}
	return true
}

func lineReflection(points [][]int) (bool, int) {
	st := make(map[[2]int]int)
	minX, maxX := points[0][0], points[0][0]
	for i := range points {
		st[[2]int{points[i][0], points[i][1]}]++
		minX = min(minX, points[i][0])
		maxX = max(maxX, points[i][0])
	}

	lineX := minX + maxX

	for i := range st {
		need := [2]int{lineX - i[0], i[1]}
		if st[i] != st[need] {
			return false, 0
		}
	}

	return true, lineX
}

func minDistanceXY(s string) int {
	findX := true
	l := 0
	for l = range s {
		if s[l] == 'X' {
			break
		} else if s[l] == 'Y' {
			findX = false
			break
		}
	}
	minD := len(s)
	for r := l + 1; r < len(s); r++ {
		if s[r] == 'X' && !findX {
			minD = min(minD, r-l)
			findX = true
		} else if s[r] == 'X' && findX {
			l = r
		} else if s[r] == 'Y' && !findX {
			l = r
		} else if s[r] == 'Y' && findX {
			minD = min(minD, r-l)
			findX = false
		}
	}

	return minD
}
