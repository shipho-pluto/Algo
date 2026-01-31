package main

import (
	"fmt"
)

type statistic struct {
	userID int
	steps  int
}

type result struct {
	userIDs []int
	steps   int
}

func getChampions(stat [][]statistic) result {
	st := make(map[int]int)
	n := len(stat)
	for j := 0; j < len(stat[0]); j++ {
		st[stat[0][j].userID] = stat[0][j].steps
	}
	for i := 1; i < n; i++ {
		nst := make(map[int]int)
		for j := 0; j < len(stat[i]); j++ {
			nst[stat[i][j].userID] = stat[i][j].steps
		}
		for j := 0; j < len(stat[0]); j++ {
			if v, ok := nst[stat[0][j].userID]; ok {
				st[stat[0][j].userID] += v
			} else {
				delete(st, stat[0][j].userID)
			}
		}
	}

	maxStep := 0
	for _, v := range st {
		maxStep = max(maxStep, v)
	}
	res := result{[]int{}, maxStep}
	for i := range st {
		if st[i] == maxStep {
			res.userIDs = append(res.userIDs, i)
		}
	}
	return res
}

func main() {
	stat := [][]statistic{
		{{1, 2000}, {2, 1500}}, {{2, 1000}},
	}
	stat2 := [][]statistic{
		{{1, 1000}, {2, 1500}, {3, 1500}},
		{{2, 1000}, {1, 500}, {3, 1000}},
	}
	fmt.Println(getChampions(stat))
	fmt.Println(getChampions(stat2))
}
