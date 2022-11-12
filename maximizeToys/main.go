package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 12, 5, 111, 200, 1000, 10}
	N := len(arr)
	K := 50

	c := toyCount(N, K, arr)
	fmt.Println(c)
}

func toyCount(c, j int, b []int) int {
	var count int
	sort.Ints(b)
	fmt.Println(b)
	for i := 0; i < j; i++ {
		if b[i] <= j {
			count++
			j = j - b[i]
		} else {
			break
		}
	}
	return count
}
