package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		m   int
		buf []string
	)
	fmt.Scanln(&m)
	m++
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for ; m > 0; m-- {
		scanner.Scan()
		buf = strings.Split(scanner.Text(), " ")
		mid(StrToInt(buf[0]), StrToInt(buf[1]))
	}
}
func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func mid(n, t int) {
	if t == 1 {
		fmt.Println(0)
		return
	}
	p := 1
	for i := 2; i < min(n, t)+1; i++ {
		p = (p * i) % t
	}
	fmt.Println(p)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
