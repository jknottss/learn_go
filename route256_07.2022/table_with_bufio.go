package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var (
		tables   int
		lines    int
		columns  int
		sortcol  int
		itersort int
	)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &tables)
	for k := 0; k < tables; k++ {
		fmt.Fscan(in, &lines, &columns)
		table := make([][]int, lines)
		for i := range table {
			table[i] = make([]int, columns)
		}
		for i := 0; i < lines; i++ {
			for j := 0; j < columns; j++ {
				fmt.Fscan(in, &table[i][j])
			}
		}
		fmt.Fscan(in, &itersort)
		for i := 0; i < itersort; i++ {
			fmt.Fscan(in, &sortcol)
			sort.SliceStable(table, func(i, j int) bool { return table[i][sortcol-1] < table[j][sortcol-1] })
		}
		for _, val := range table {
			for _, nbr := range val {
				fmt.Fprint(out, nbr, " ")
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}
