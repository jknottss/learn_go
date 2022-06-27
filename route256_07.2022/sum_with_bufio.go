package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int

	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var n, res int
		nabor := make(map[int]int)
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			var m int
			fmt.Fscan(in, &m)
			nabor[m] += 1
		}
		for i, value := range nabor {
			res += (value / 3 * 2 * i) + (value % 3 * i)
		}
		fmt.Fprintln(out, res)
	}
}
