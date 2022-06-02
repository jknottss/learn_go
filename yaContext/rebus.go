package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {

	alph := map[int]string{
		0:  "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
		16: "q",
		17: "r",
		18: "s",
		19: "t",
		20: "u",
		21: "v",
		22: "w",
		23: "x",
		24: "y",
		25: "z",
		26: " ",
	}
	var n, input, vrem string

	fmt.Scan(&n)

	count, _ := strconv.Atoi(n)
	for i := 0; i < count; i++ {
		fmt.Scan(&input)
		vrem = vrem + " " + input
	}
	log.Println(vrem)
	split := strings.Split(vrem, " ")
	log.Println(split)
	for i := 0; i <= count; i++ {
		nbr, _ := strconv.Atoi(split[i])
		if i > 0 {
			tmp, _ := strconv.Atoi(split[i-1])
			nbr = nbr - tmp
			if nbr < 0 {
				nbr *= -1
			}
			log.Print(nbr)
		}
		nbr = int(math.Log2(float64(nbr)))
		fmt.Print(alph[nbr])
	}
}
