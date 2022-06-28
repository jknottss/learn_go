package main

import (
	"fmt"
	"sort"
)

func main() {
	var tables int
	var lines int
	var columns int
	var sortcol int
	var itersort int
	fmt.Scan(&tables)
	for k := 0; k < tables; k++ {
		fmt.Scan(&lines, &columns)
		table := make([][]int, lines)
		for i := range table {
			table[i] = make([]int, columns)
		}
		for i := 0; i < lines; i++ {
			for j := 0; j < columns; j++ {
				fmt.Scan(&table[i][j])
			}
		}
		fmt.Scan(&itersort)
		for i := 0; i < itersort; i++ {
			fmt.Scan(&sortcol)
			sort.SliceStable(table, func(i, j int) bool { return table[i][sortcol-1] < table[j][sortcol-1] })
		}
		for _, val := range table {
			for _, nbr := range val {
				fmt.Print(nbr, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
