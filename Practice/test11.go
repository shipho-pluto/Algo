package main

import "fmt"

type CMap11 struct {
	data         map[[26]int]int
	defaultValue int
}

func (cm *CMap11) Get(key [26]int) int {
	if value, exists := cm.data[key]; exists {
		return value
	}
	return cm.defaultValue
}

func (cm *CMap11) Set(key [26]int, value int) {
	cm.data[key] = value
}

func Minus11(arr1, arr2 [26]int) [26]int {
	a := [26]int{}
	for i := 0; i < 26; i++ {
		a[i] = arr1[i] - arr2[i]
	}
	return a
}

func StringInString(str, pres string) int { // O(m + n)*26
	mask := [26]int{}
	m, n := len(str), len(pres)
	for i := 0; i < n; i++ {
		mask[int(pres[i]-'a')] += 1
	}
	sum := [26]int{}
	maskMap := CMap11{make(map[[26]int]int), -2}
	maskMap.Set([26]int{}, -1)
	for i := 0; i < m; i++ {
		sum[str[i]-'a'] += 1
		if maskMap.Get(Minus11(sum, mask)) != -2 {
			return maskMap.Get(Minus11(sum, mask)) + 1
		}
		maskMap.Set(sum, i)
	}
	return -1
}

func sts(str, pod string) int { // O(m) + 26
	maskP := [26]int{}
	maskW := [26]int{}
	n := len(pod)
	for i := 0; i < n; i++ {
		maskP[int(pod[i]-'a')] += 1
		maskW[int(str[i]-'a')] += 1
	}
	eq := 0
	for i := 0; i < 26; i++ {
		if maskP[i] == maskW[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	m := len(str)
	for i := n; i < m; i++ {
		prev := int(str[i] - 'a')
		if maskP[prev] == maskW[prev] {
			eq--
		}
		maskW[prev] += 1
		if maskP[prev] == maskW[prev] {
			eq++
		}

		post := int(str[i-n] - 'a')
		if maskP[post] == maskW[post] {
			eq--
		}
		maskW[post] -= 1
		if maskP[post] == maskW[post] {
			eq++
		}

		if eq == 26 {
			return i - n + 1
		}
	}
	return -1
}

func main() {
	a := "babbsda"
	p := "bab"
	fmt.Println(StringInString(a, p))
	fmt.Println(sts(a, p))
}
