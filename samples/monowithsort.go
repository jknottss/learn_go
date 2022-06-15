package main

import (
	"fmt"
	"sort"
)

func ismono(sl []int) bool {
	up := sort.SliceIsSorted(sl, func(i, j int) bool { return sl[i] < sl[j] })
	down := sort.SliceIsSorted(sl, func(i, j int) bool { return sl[i] > sl[j] })

	if len(sl) < 2 || up || down {
		return true
	}

	return false
}

func main() {
	fmt.Println(ismono([]int{1, 2, 3}))
	fmt.Println(ismono([]int{5, 10, 1}))
	fmt.Println(ismono([]int{1, 1, 1}))
}
