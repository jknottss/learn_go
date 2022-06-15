package main

import "fmt"

func test(sl []int) bool {
	var direction bool
	if len(sl) < 2 {
		return true
	}

	direction = sl[0] < sl[1]

	for i := 1; i < len(sl)-1; i++ {
		if sl[i] > sl[i+1] && direction {
			return false
		} else if sl[i] < sl[i+1] && !direction {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(test([]int{1, 2, 3}))
	fmt.Println(test([]int{5, 6, 4}))
	fmt.Println(test([]int{1, 1, 1}))
}
