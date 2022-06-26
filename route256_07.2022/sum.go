package main

import "fmt"

func main() {
	var in, sum int

	fmt.Scanln(&in)
	for i := 0; i < in; i++ {
		var numbers int
		fmt.Scan(&numbers)
		myMap := make(map[int]int, numbers)
		for j := 0; j < numbers; j++ {
			var nbr int
			fmt.Scan(&nbr)
			myMap[nbr]++
		}
		for num, count := range myMap {
			switch {
			case count < 3:
				sum += num * count
			case count == 3:
				sum += num * 2
			case count > 3 && count%3 != 0:
				sum += num * (count - count/3)
			case count > 3 && count%3 == 0:
				sum += num * 2 * count / 3
			}
		}
		fmt.Println(sum)
		sum = 0
	}
}
